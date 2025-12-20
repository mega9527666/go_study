package main

import (
	"mega/account_server/account_reqhandler"
	"mega/common/config"
	"mega/common/db_config"
	"mega/engine/logger"
	"mega/engine/string_util"
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

	dataObj := map[string]interface{}{
		"account": "abcd",
		"chan":    66,
	}

	// result := dataObj["chan"]
	result, ok := string_util.GetIntFromMap(dataObj, "chan")
	logger.Log("result--=", result, ok)
	result1 := dataObj["account"]
	logger.Log("result--account=", result1)
	// accountModel, err := account_model.GetAccountByAccount("ab1d")
	// if err != nil {
	// 	logger.Warn("查找账号失败===", accountModel)
	// } else {
	// 	if accountModel != nil {
	// 		if accountModel.Pass == "12123165" {
	// 			logger.Warn("密码相同===", accountModel)
	// 		} else {
	// 			logger.Warn("密码不同===", accountModel)
	// 			// http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrPasswordWrong})
	// 		}
	// 	} else {
	// 		logger.Warn("账号不存在===", accountModel, err)
	// 	}
	// }

	account_reqhandler.ListenAndServe(port)
}
