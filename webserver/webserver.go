package main

import (
	"mega/engine/logger"
	"mega/webserver/webreqhandler"
	// "github.com/gorilla/websocket"
)

func main() {
	logger.Log("webserver.main")
	webreqhandler.Init(9090)

}
