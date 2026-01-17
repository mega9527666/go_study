package main

import (
	"mega/account_server/account_reqhandler"
	"mega/common/config"
	"mega/engine/logger"
	"mega/engine/string_util"
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
	logger.Info("account_server.main", os.Args, env, http_port)

	dataObj := map[string]interface{}{
		"account": "abcd",
		"chan":    66,
	}

	// result := dataObj["chan"]
	result, ok := string_util.GetIntFromMap(dataObj, "chan")
	logger.Log("result--=", result, ok)
	result1 := dataObj["account"]
	logger.Log("result--account=", result1)
	// accountModel, err := account_model.GetAccountByAccount("abcd")
	// if err != nil {
	// 	logger.Warn("查找账号失败===", accountModel)
	// } else {
	// 	if accountModel != nil {
	// 		if accountModel.Pass == "123456" {
	// 			logger.Warn("密码相同===", accountModel)
	// 			var token string = md5_helper.CreateToken(accountModel.Account)
	// 			accountModel.NickName = accountModel.Account
	// 			accountModel.Token = token
	// 			rows, err := account_model.UpdateAccount(accountModel)
	// 			if err != nil {
	// 			}
	// 		} else {
	// 			logger.Warn("密码不同===", accountModel)
	// 			// http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrPasswordWrong})
	// 		}
	// 	} else {
	// 		logger.Warn("账号不存在===", accountModel, err)
	// 	}
	// }

	account_reqhandler.ListenAndServe(http_port)
}
