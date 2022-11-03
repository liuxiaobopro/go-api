package db

import (
	"embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/liuxiaobopro/go-api/config"
	"github.com/liuxiaobopro/go-api/global"

	_ "github.com/go-sql-driver/mysql"
	"github.com/liuxiaobopro/go-lib/console"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var (
	Mysql *xorm.EngineGroup
)

func InitDb(FsMysql embed.FS) {
	var err error
	MDns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		global.Conf.Database.Mdb.Username,
		global.Conf.Database.Mdb.Password,
		global.Conf.Database.Mdb.Host,
		global.Conf.Database.Mdb.Port,
		global.Conf.Database.Mdb.Database,
		global.Conf.Database.Mdb.Charset)
	SDns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local",
		global.Conf.Database.Sdb.Username,
		global.Conf.Database.Sdb.Password,
		global.Conf.Database.Sdb.Host,
		global.Conf.Database.Sdb.Port,
		global.Conf.Database.Sdb.Database,
		global.Conf.Database.Sdb.Charset)

	Mysql, err = xorm.NewEngineGroup("mysql", []string{MDns, SDns})
	if err != nil {
		console.Console.Error("初始化数据库失败", err.Error())
	}
	// 连接池参数:空闲数、最大连接数、连接最大生存时间
	Mysql.SetMaxIdleConns(10)
	Mysql.SetMaxOpenConns(100)
	Mysql.SetConnMaxLifetime(3 * time.Hour)

	Mysql.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, global.Conf.Database.Mdb.Prefix))
	// Mysql.SetColumnMapper(core.SnakeMapper{})
	Mysql.SetColumnMapper(core.GonicMapper{})

	if global.Conf.App.SqlSaveToFile {
		// 获取项目根路径
		pwd, _ := os.Getwd()
		sqllog := path.Join(pwd, global.Conf.App.SqlLogPath, "sql.log")
		// 判断是否存在
		if _, err := os.Stat(sqllog); os.IsNotExist(err) {
			// 创建文件夹
			os.MkdirAll(filepath.Dir(sqllog), os.ModePerm)
			// 创建文件
			os.Create(sqllog)
		}
		f, err := os.Create(sqllog)
		if err != nil {
			console.Console.Error("创建sql.log文件失败", err.Error())
			return
		}
		Mysql.SetLogger(log.NewSimpleLogger(f)) // 设置 logger
	}

	Mysql.SetLogLevel(log.LOG_DEBUG)
	Mysql.ShowSQL(true)

	if err := Mysql.Ping(); err != nil {
		console.Console.Error("连接数据库失败", err.Error())
	}

	// 同步表
	if err := Mysql.Sync2(config.Tables()...); err != nil {
		console.Console.Error("同步数据库表失败", err.Error())
	}
	global.Db = Mysql

	if !config.Conf.App.IsInstall {
		runSql(Mysql, FsMysql)
	}
}

// 执行初始化sql
func runSql(Mysql *xorm.EngineGroup, FsMysql embed.FS) {
	// // 获取同目录mysql.sql文件内容
	// pwd, _ := os.Getwd()
	// filepath := filepath.Join(pwd, "/core/db/mysql.sql")
	// if _, err := os.Stat(filepath); os.IsNotExist(err) {
	// 	console.Console.Error("mysql.sql文件不存在", err.Error())
	// }
	// sql, err := ioutil.ReadFile(filepath)

	sql, err := FsMysql.ReadFile("core/db/mysql.sql")
	if err != nil {
		console.Console.Error("读取mysql.sql文件失败", err.Error())
	}
	s := string(sql)
	// 按照分号分割
	sqls := strings.Split(s, ";")
	// 过滤空字符
	for _, sql := range sqls {
		// 多删几遍
		for i := 0; i < 2; i++ {
			// 删除头尾的换行符
			sql = strings.Trim(sql, "\r\n")
			// 删除头尾的空格
			sql = strings.TrimSpace(sql)
		}
		if sql != "" {
			_, err := Mysql.Exec(sql)
			if err != nil {
				console.Console.Error("执行sql语句失败", err.Error())
			}
		}
	}
}
