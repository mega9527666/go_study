package webreqhandler

import (
	"encoding/json"
	"mega/engine/logger"
	"net/http"
	"os"
	"strconv"
)

type httpHandleFunc func(w http.ResponseWriter, r *http.Request)

var routesMap = map[string]httpHandleFunc{
	"/mega": dispatcher(megaHandler),
	"/abcd": dispatcher(abcdHandler),
}

func init() {
	logger.Log("webreqhandler.init")
}

func ListenAndServe(port int) {
	var portStr string = strconv.Itoa(port)
	var dirPath string = "public/public" + portStr
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			logger.Log("创建文件夹失败==", dirPath, err)
		} else {
			logger.Log("创建文件夹成功==", dirPath)
		}
	} else {
		logger.Log("已经存在文件夹==", dirPath)
	}

	// 设置静态文件服务器的根目录为当前目录
	fs := http.FileServer(http.Dir(dirPath))
	// 设置路由规则，将所有请求重定向到静态文件服务器
	http.Handle("/", fs)

	/*使用键输出地图值 */
	for route := range routesMap {
		logger.Log("routesMap=======", route, routesMap[route])
		http.HandleFunc(route, routesMap[route])
	}
	http.ListenAndServe(":"+portStr, nil)
}

// 通用的分发函数（中间件）
func dispatcher(next httpHandleFunc) httpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// go indexHandler(w, r, next)
		indexHandler(w, r, next)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, next httpHandleFunc) {
	// 在执行实际的请求处理之前做一些处理
	logger.Log("通用分发函数：请求到来，执行前处理...", r.RequestURI, r.Host, r.RemoteAddr)
	// 你可以在这里加入公共的处理逻辑，例如验证、日志记录等
	// 调用下一个处理函数
	next(w, r)
	// 在实际的请求处理之后做一些处理
	logger.Log("通用分发函数：请求处理完毕，执行后处理...", r.RequestURI, r.Host, r.RemoteAddr)
}

func megaHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("megaHandler=Host=", r.RequestURI, r.Host, r.RemoteAddr)
}

type Response struct {
	Message string `json:"message"`
}

func abcdHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("abcdHandler==", r.RequestURI, r.Host, r.RemoteAddr)
	logger.Log("abcdHandler=param=Body", r.Body)
	var queryMap = r.URL.Query()
	logger.Log("abcdHandler=param=queryMap", queryMap)
	logger.Log("abcdHandler=param=gameId", queryMap.Get("gameId"))
	// 设置响应的 Content-Type 为 text/plain
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/plain")
	// w.Header().Set("Content-Type", "application/json")
	// 创建一个响应对象
	response := Response{Message: "Hello, JSON abcd!"}
	// 向客户端写入响应内容
	// 将结构体编码为 JSON 并写入响应体
	err := json.NewEncoder(w).Encode(response)
	// _, err := w.Write([]byte("Hello, abcd!"))
	if err != nil {
		// http.Error(w, "Unable to encode JSON", http.StatusInternalServerError)
		logger.Error("abcdHandler=error=", err)
	}
	// w.WriteHeader(http.StatusOK)
	// w.Write()
}
