package account_reqhandler

import (
	"mega/common/model/account_model"
	"mega/engine/error_code"
	"mega/engine/http_common"
	"mega/engine/logger"
	"mega/engine/string_util"
	"net/http"
)

var routesMap = map[string]http_common.HttpHandleFunc{
	"/register": http_common.Dispatcher(register),
	"/login":    http_common.Dispatcher(login),
}

func ListenAndServe(port int) {
	http_common.ListenAndServe(port, routesMap)
}

func register(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{}) {
	logger.Log("register=", ip, dataObj)
	account, ok := string_util.GetStringFromMap(dataObj, "account")
	if !ok {
		http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrParam})
		return
	}
	logger.Log("register=====account", account)
	pass, ok := string_util.GetStringFromMap(dataObj, "pass")
	if !ok {
		http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrParam})
		return
	}
	logger.Log("register=====pass", pass)

	var accountModel account_model.Account = account_model.Account{
		Account: account,
		Pass:    pass,
	}
	logger.Log("开始查询...")
	account_model.IsAccountExist(accountModel.Account, func(exists bool, err error) {
		if err != nil {
			logger.Log("查询错误: ", err)
			http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrInternal})
		} else {
			logger.Log("账号存在: ", exists)
			if !exists {
				account_model.InsertAccount_callback(&accountModel, func(i int64, err error) {
					logger.Log("InsertAccount_callback======", i, err)
					accountModel.ID = i
					var respModel http_common.HttpResponseModel = http_common.HttpResponseModel{Code: error_code.OK}
					respModel.Data = map[string]interface{}{
						"account": account,
					}
					http_common.SendHttpResponseModel(w, respModel)
				})
			} else {
				http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrAccountExist})
			}
		}
	})

	// var respModel http_common.HttpResponseModel = http_common.HttpResponseModel{Code: error_code.OK}
	// respModel.Data = map[string]interface{}{
	// 	"need_hotupdate": false, //需要热更新
	// 	"force_update":   false,
	// 	"ip":             ip,
	// }
	// http_common.SendHttpResponseModel(w, respModel)
}

func login(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{}) {
	logger.Log("register=", ip, dataObj)
	account, ok := string_util.GetStringFromMap(dataObj, "account")
	if !ok {
		http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrParam})
		return
	}
	logger.Log("register=====account", account)
	// pass, ok := string_util.GetStringFromMap(dataObj, "pass")
	// if !ok {
	// 	http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrParam})
	// 	return
	// }
}
