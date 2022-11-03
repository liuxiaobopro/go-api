package main

import (
	"embed"

	coreConfig "github.com/liuxiaobopro/go-api/core/config"
	"github.com/liuxiaobopro/go-api/core/db"
	"github.com/liuxiaobopro/go-api/core/gin"
	"github.com/liuxiaobopro/go-api/core/upload/qiniu"
)

//go:embed "config/yaml/*"
var FS embed.FS

//go:embed "core/db/mysql.sql"
var FsMysql embed.FS

func main() {
	// 初始化配置
	coreConfig.InitConfig(FS)
	// 初始化数据库
	db.InitDb(FsMysql)
	// 初始化七牛云
	qiniu.InitQiniu()
	// 初始化gin
	gin.InitGin()
}
