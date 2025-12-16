package main

import (
	"mega/account_server/account_reqhandler"
	"mega/common/config"
	"mega/common/db_config"
	"mega/common/model/account_model"
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
	db_config.InitDb(config.Environment)
	config.ServerType = config.ServerType_List.Account_server
	logger.Info("account_server.main", os.Args, env, port)

	// var accountModel account_model.Account = account_model.Account{
	// 	Account: "abcd",
	// 	Pass:    "123456",
	// }

	// logger.Log("查询已启动，继续执行其他任务...")
	// // 等待回调完成（实际项目中不需要这样）
	// time.Sleep(1 * time.Second)
	accountModel, err := account_model.GetAccountByAccount("ab1d")
	if err != nil {
		logger.Warn("查找账号失败===", accountModel)
	} else {
		if accountModel != nil {
			if accountModel.Pass == "12123165" {
				logger.Warn("密码相同===", accountModel)
			} else {
				logger.Warn("密码不同===", accountModel)
				// http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrPasswordWrong})
			}
		} else {
			logger.Warn("账号不存在===", accountModel, err)
		}
	}

	account_reqhandler.ListenAndServe(port)
}
