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
	exists, err := account_model.IsAccountExist(account)
	if err != nil {
		http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrInternal})
	} else {
		if !exists {
			id, err := account_model.InsertAccount(&accountModel)
			if err != nil {
				http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrInternal})
			} else {
				accountModel.ID = id
				var respModel http_common.HttpResponseModel = http_common.HttpResponseModel{Code: error_code.OK}
				respModel.Data = map[string]interface{}{
					"account": account,
				}
				http_common.SendHttpResponseModel(w, respModel)
			}
		} else {
			http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrAccountExist})
		}
	}
}

func login(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{}) {
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

	accountModel, err := account_model.GetAccountByAccount(account)
	if err != nil {
		logger.Warn("查找账号失败===", accountModel)
		http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrInternal})
	} else {
		if accountModel != nil {
			if accountModel.Pass == pass {

			} else {
				http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrPasswordWrong})
			}
		} else {
			// logger.Warn("账号不存在===", accountModel, err)
			http_common.SendHttpResponseModel(w, http_common.HttpResponseModel{Code: error_code.ErrAccountNotFound})
		}
	}
}
