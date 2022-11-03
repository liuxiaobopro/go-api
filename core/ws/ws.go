package ws

import (
	"github.com/liuxiaobopro/go-api/global"

	"github.com/gin-gonic/gin"
)

func InitWs(g *gin.Engine) {
	go WebsocketManager.Start()
	go WebsocketManager.SendService()
	go WebsocketManager.SendService()
	go WebsocketManager.SendGroupService()
	go WebsocketManager.SendGroupService()
	go WebsocketManager.SendAllService()
	go WebsocketManager.SendAllService()

	// go TestSendGroup()
	// go TestSendAll()

	g.GET(global.Conf.Ws.Route, WebsocketManager.WsClient)
}
