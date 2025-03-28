package main

import (
	"mega/engine/logger"
	"mega/webserver/webreqhandler"
	"os"
	"strconv"
	// "github.com/gorilla/websocket"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Error("初始化端口失败:", os.Args, err)
		return
	}
	logger.Info("webserver.main", port)
	webreqhandler.ListenAndServe(port)

}
