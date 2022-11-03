package ws

import (
	wsCore "github.com/liuxiaobopro/go-api/core/ws"
)

func SendAll(msg string) {
	wsCore.WebsocketManager.SendAll([]byte(msg))
}

func SendGroup(group, msg string) {
	wsCore.WebsocketManager.SendGroup(group, []byte(msg))
}
