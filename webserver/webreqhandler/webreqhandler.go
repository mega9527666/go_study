package webreqhandler

import (
	"mega/engine/error_code"
	"mega/engine/http_common"
	"mega/engine/logger"
	"net/http"
)

var routesMap = map[string]http_common.HttpHandleFunc{
	// "/mega":        http_common.Dispatcher(megaHandler),
	"/init_client": http_common.Dispatcher(init_client_handler),
}

func init() {
	// logger.Log("webreqhandler.init")
}

func ListenAndServe(port int) {
	http_common.ListenAndServe(port, routesMap)
}

func megaHandler(w http.ResponseWriter, r *http.Request, ip, dataObj map[string]interface{}) {
	logger.Log("megaHandler=Host=", ip)
}

type Response struct {
	Message string `json:"message"`
}

// 定义示例结构体
// 根据你的实际需求调整字段类型和标签
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func init_client_handler(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{}) {
	logger.Log("init_client_handler=param=Body", ip)
	var respModel http_common.HttpResponseModel = http_common.HttpResponseModel{Code: error_code.OK}
	respModel.Data = map[string]interface{}{
		"need_hotupdate": false, //需要热更新
	}
	http_common.SendHttpResponseModel(w, respModel)
}
