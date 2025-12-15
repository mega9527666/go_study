package main

import (
	"mega/account_server/account_reqhandler"
	"mega/common/config"
	dbconfig "mega/common/db_config"
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
	env, err := strconv.Atoi(os.Args[2])
	if err != nil {
		logger.Error("初始化环境失败:", os.Args, err)
		return
	}

	config.Environment = env
	dbconfig.InitDb(config.Environment)
	config.ServerType = config.ServerType_List.Account_server
	logger.Info("account_server.main", os.Args, env, port)

	account_reqhandler.ListenAndServe(port)
}
