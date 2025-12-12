package main

import (
	"mega/account_server/account_reqhandler"
	"mega/engine/logger"
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
	logger.Info("account_server.main", port)
	account_reqhandler.ListenAndServe(port)
}
