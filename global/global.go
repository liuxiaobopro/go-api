package global

import (
	"github.com/liuxiaobopro/go-api/config"

	"github.com/gorilla/websocket"
	"xorm.io/xorm"
)

var (
	Conf          *config.Config    // 配置
	Db            *xorm.EngineGroup // 数据库
	WebsocketConn *websocket.Conn   // websocket连接
)
