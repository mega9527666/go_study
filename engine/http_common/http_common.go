package http_common

import (
	"encoding/json"
	"io"
	Md5Helper "mega/engine"
	"mega/engine/logger"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type HttpHandleFunc func(w http.ResponseWriter, r *http.Request)
type HttpCustomHandleFunc func(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{})

func ListenAndServe(port int, routesMap map[string]HttpHandleFunc) {
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
		// logger.Log("已经存在文件夹==", dirPath)
	}

	// 设置静态文件服务器的根目录为当前目录
	fs := http.FileServer(http.Dir(dirPath))
	// 设置路由规则，将所有请求重定向到静态文件服务器
	http.Handle("/", fs)

	/*使用键输出地图值 */
	for route := range routesMap {
		// logger.Log("routesMap=======", route, routesMap[route])
		http.HandleFunc(route, routesMap[route])
	}
	http.ListenAndServe(":"+portStr, nil)
}

// 通用的分发函数（中间件）
func Dispatcher(next HttpCustomHandleFunc) HttpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许所有来源，生产环境可改为指定域名
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头，这里加上你需要的自定义头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,X-token")
		// 如果你希望前端能够读取到某些响应头，则用 Expose-Headers
		w.Header().Set("Access-Control-Expose-Headers", "X-My-Custom-Header")
		// 在执行实际的请求处理之前做一些处理

		var contentType string = r.Header.Get("content-type")
		if len(contentType) == 0 {
			r.Header.Set("content-type", "application/json")
		}
		logger.Log("dispatcher=param=Content-Type", r.Header.Get("content-type"))
		commonHandler(w, r, next)
	}
}

func commonHandler(w http.ResponseWriter, r *http.Request, next HttpCustomHandleFunc) {
	var ip string = GetClientIP(r)
	logger.Log("通用分发函数：请求到来，执行前处理...", ip)
	// r.Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "commonHandler read body error", http.StatusBadRequest)
		return
	}
	logger.Log("通用分发函数：body", body)
	logger.Log("通用分发函数：body2", string(body))

	// 1) 解析 form
	datas, err := url.ParseQuery(string(body))
	if err != nil {
		logger.Error("commonHandler error=", err)
		return
	}
	logger.Log("commonHandler datas type", datas)
	dataStr := datas["data"][0]
	k := datas["k"][0]

	var dataObj map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &dataObj); err != nil {
		// panic(err)
		logger.Error("commonHandler error=", err)
		return
	}

	// fmt.Printf("data type-%T\n", datas["data"])
	// fmt.Printf("data j type %T\n", datas["k"])
	logger.Log("commonHandler datas", datas)
	logger.Log("commonHandler dataStr", dataStr)
	logger.Log("commonHandler k", k)
	logger.Log("commonHandler dataObj", dataObj)
	logger.Log("commonHandler channel", dataObj["channel"])
	logger.Log("commonHandler t", dataObj["t"])
	logger.Log("commonHandler v", dataObj["v"])

	var dataK string = Md5Helper.GetMd5_default(dataStr)
	logger.Log("commonHandler checkKey=", dataK, k == dataK)
	// let
	// fmt.Println("Body:", string(body))
	// 你可以在这里加入公共的处理逻辑，例如验证、日志记录等
	// 调用下一个处理函数
	next(w, r, ip, dataObj)
	// 在实际的请求处理之后做一些处理
	logger.Log("通用分发函数：请求处理完毕，执行后处理...", r.RequestURI, r.Host, r.RemoteAddr)
}

func GetClientIP(r *http.Request) string {
	// 1. X-Forwarded-For
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// 可能是多IP: "client, proxy1, proxy2"
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	// 2. X-Real-IP
	xrp := r.Header.Get("X-Real-Ip")
	if xrp != "" {
		return xrp
	}

	// 3. RemoteAddr（兜底）
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
