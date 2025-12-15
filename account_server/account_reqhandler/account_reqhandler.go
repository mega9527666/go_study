package account_reqhandler

import (
	"mega/engine/error_code"
	"mega/engine/http_common"
	"mega/engine/logger"
	"mega/engine/string_util"
	"net/http"
)

var routesMap = map[string]http_common.HttpHandleFunc{
	"/register": http_common.Dispatcher(register),
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
	// var respModel http_common.HttpResponseModel = http_common.HttpResponseModel{Code: error_code.OK}
	// respModel.Data = map[string]interface{}{
	// 	"need_hotupdate": false, //需要热更新
	// 	"force_update":   false,
	// 	"ip":             ip,
	// }
	// http_common.SendHttpResponseModel(w, respModel)
}
