package account_reqhandler

import (
	"mega/engine/http_common"
	"mega/engine/logger"
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
	// var respModel http_common.HttpResponseModel = http_common.HttpResponseModel{Code: error_code.OK}
	// respModel.Data = map[string]interface{}{
	// 	"need_hotupdate": false, //需要热更新
	// 	"force_update":   false,
	// 	"ip":             ip,
	// }
	// http_common.SendHttpResponseModel(w, respModel)
}
