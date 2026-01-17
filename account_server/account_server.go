package main

import (
	"mega/account_server/account_reqhandler"
	"mega/common/config"
	"mega/engine/logger"
	"mega/grpc/grpc_server"
	"os"
	"strconv"
	// "github.com/gorilla/websocket"
)

func main() {
	http_port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Error("初始化端口失败:", os.Args, err)
		return
	}
	env, err := strconv.Atoi(os.Args[2])
	if err != nil {
		logger.Error("初始化环境失败:", os.Args, err)
		return
	}

	config.Env = env
	config.ServerType = config.ServerType_List.Account_server
	//读取yaml配置文件
	config.InitConfig(http_port)

	//启动grpc服务
	go func() {
		err := grpc_server.StartGrpcServer(
			config.Now_ServerItem.GrpcPort,
		)
		if err != nil {
			logger.Warn("grpc start error:", err)
		}
	}()

	account_reqhandler.ListenAndServe(http_port)
}
