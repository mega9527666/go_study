package webreqhandler

import (
	"mega/engine/logger"
	"net/http"
	"os"
	"strconv"
)

type httpHandleFunc func(w http.ResponseWriter, r *http.Request)

// 通用的分发函数（中间件）
func dispatcher(next httpHandleFunc) httpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		go indexHandler(w, r, next)
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

func abcdHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("abcdHandler=Host=", r.RequestURI, r.Host, r.RemoteAddr)
}

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
