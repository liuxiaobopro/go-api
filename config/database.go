package config

import (
	"github.com/liuxiaobopro/go-api/cmd/reverse/models"
)

type Database struct {
	Mdb struct { // 主库
		Host     string // 数据库地址
		Port     int    // 数据库端口
		Database string // 数据库名称
		Username string // 数据库用户名
		Password string // 数据库密码
		Charset  string // 数据库编码
		Prefix   string // 数据库表前缀
	}
	Sdb struct { // 从库
		Host     string // 数据库地址
		Port     int    // 数据库端口
		Database string // 数据库名称
		Username string // 数据库用户名
		Password string // 数据库密码
		Charset  string // 数据库编码
		Prefix   string // 数据库表前缀
	}
}

// Tables 要同步的表
func Tables() []interface{} {
	return []interface{}{
		new(models.GaCabinetOpenOperation),
		new(models.GaCabinet),
		new(models.GaMail),
		new(models.GaReceiveDep),
		new(models.GaToken),
		new(models.GaUser),
	}
}
