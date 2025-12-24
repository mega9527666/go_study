package main

import (
	"mega/common/config"
	"mega/common/db_config"
	"mega/engine/logger"
	"mega/engine/socket_common"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Error("初始化端口失败:", os.Args, err)
		return
	}
	env, err := strconv.Atoi(os.Args[2])
	if err != nil {
		logger.Error("初始化环境失败:", os.Args, err)
		return
	}

	config.Environment = env
	db_config.InitDb(config.Environment)
	config.ServerType = config.ServerType_List.Hall_server
	logger.Info("Hall_server.main", os.Args, env, port)

	// http.HandleFunc("/ws", wsHandler)
	// http.HandleFunc("/ws", socket_common.WsHandler)
	http.HandleFunc("/", socket_common.WsHandler)
	// log.Println("WebSocket 服务启动: ws://localhost:8080/ws")
	logger.Log("WebSocket 服务启动: ws://127.0.0.1:" + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}
