package main

import (
	"mega/account_server/account_reqhandler"
	"mega/common/config"
	"mega/common/db_config"
	"mega/common/model/account_model"
	"mega/engine/logger"
	"mega/engine/mysql_client"
	"mega/engine/mysql_manager"
	"os"
	"strconv"
	"time"
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
	db_config.InitDb(config.Environment)
	config.ServerType = config.ServerType_List.Account_server
	logger.Info("account_server.main", os.Args, env, port)

	logger.Log("开始查询...")
	var client *mysql_client.Db_client = mysql_manager.GetDb(db_config.Db_account, db_config.NowDbType)

	account_model.IsAccountExist(client, "abcd", func(exists bool, err error) {
		if err != nil {
			logger.Log("查询错误: ", err)
		} else {
			logger.Log("账号存在: ", exists)
		}
	})
	logger.Log("查询已启动，继续执行其他任务...")
	// 等待回调完成（实际项目中不需要这样）
	time.Sleep(1 * time.Second)

	account_reqhandler.ListenAndServe(port)
}
