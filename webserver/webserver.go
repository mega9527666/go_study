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
	// env, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// 	logger.Error("初始化环境失败:", os.Args, err)
	// 	return
	// }

	// config.Environment = env
	logger.Info("webserver.main", os.Args, port)
	webreqhandler.ListenAndServe(port)

}
