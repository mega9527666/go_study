package main

import (
	"mega/common/config"
	"mega/engine/logger"
	"mega/webserver/webreqhandler"
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
	config.ServerType = config.ServerType_List.Web_server
	//读取yaml配置文件
	config.InitConfig(http_port)

	logger.Info("webserver.main", os.Args, env, http_port)
	webreqhandler.ListenAndServe(http_port)

}
