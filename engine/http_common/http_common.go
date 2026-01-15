package http_common

import (
	"encoding/json"
	"io"
	"mega/engine/error_code"
	"mega/engine/logger"
	"mega/engine/md5_helper"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type HttpHandleFunc func(w http.ResponseWriter, r *http.Request)
type HttpCustomHandleFunc func(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{})

type HttpResponseModel struct {
	Code error_code.Code
	Data map[string]interface{}
}

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

func ListenAndServeTLS(port int, routesMap map[string]HttpHandleFunc, crtPath string, keyPath string) {
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
	mux := http.NewServeMux()
	// 静态文件服务
	// https://domain.com/static/xxx
	fs := http.FileServer(http.Dir("./static"))
	// 设置路由规则，将所有请求重定向到静态文件服务器
	mux.Handle("/", fs)
	/*使用键输出地图值 */
	for route := range routesMap {
		// logger.Log("routesMap=======", route, routesMap[route])
		mux.HandleFunc(route, routesMap[route])
	}
	addr := ":" + portStr
	// httpError := http.ListenAndServeTLS(
	// 	addr,
	// 	"./server.crt",
	// 	"./server.key",
	// 	mux,
	// )

	httpError := http.ListenAndServeTLS(
		addr,
		crtPath,
		keyPath,
		mux,
	)
	if err != nil {
		// log.Fatal(err)
		logger.Error("ListenAndServeTLS error", port, crtPath, keyPath, httpError)
	}
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
			contentType = "application/json"
			r.Header.Set("content-type", contentType)
		}
		logger.Log("dispatcher=param=Content-Type", contentType)
		commonHandler(w, r, next, contentType)
	}
}

func commonHandler(w http.ResponseWriter, r *http.Request, next HttpCustomHandleFunc, contentType string) {
	var ip string = GetClientIP(r)
	// logger.Log("通用分发函数：请求到来，执行前处理...", ip)
	// r.Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read body error", http.StatusBadRequest)
		return
	}
	// logger.Log("通用分发函数：body", body)
	// logger.Log("通用分发函数：body2", string(body))
	if contentType != "application/x-www-form-urlencoded" {
		logger.Error("commonHandler  不支持的contentType=", ip, contentType)
		SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrParam})
	} else {
		datas, err := url.ParseQuery(string(body))
		if err != nil {
			logger.Error("commonHandler ParseQuery error=", ip, err)
			SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrParam})
			return
		}
		logger.Log("commonHandler datas type", datas)
		// if len(datas) == 0 {
		// 	// 没有任何参数
		// 	SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrParam})
		// }

		vals, ok := datas["data"]
		if !ok {
			logger.Error("commonHandler  data 不存在 或 没有值=", ip, vals, ok)
			SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrParam})
			return
		}
		ks, ok := datas["k"]
		if !ok {
			// k 不存在 或 没有值
			logger.Error("commonHandler  k 不存在 或 没有值=", ip, ks, ok)
			SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrParam})
			return
		}

		dataStr := datas["data"][0]
		k := datas["k"][0]

		var dataObj map[string]interface{}
		if err := json.Unmarshal([]byte(dataStr), &dataObj); err != nil {
			logger.Error("commonHandler Unmarshal error=", ip, err)
			SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrParam})
			return
		}
		var dataK_encry string = md5_helper.GetMd5_encrypt(dataStr)
		// logger.Log("commonHandler checkKey=", dataK, k == dataK)
		// logger.Log("commonHandler dataK_encry=", dataK_encry, k == dataK_encry)
		if k == dataK_encry {
			next(w, r, ip, dataObj)
		} else {
			SendHttpResponseModel(w, HttpResponseModel{Code: error_code.ErrBadMd5})
		}
	}

}

func SendHttpResponseModel(w http.ResponseWriter, responseModel HttpResponseModel) {
	// 向客户端写入响应内容
	w.WriteHeader(http.StatusOK)
	// 将结构体编码为 JSON 并写入响应体
	encodeWrr := json.NewEncoder(w).Encode(responseModel)
	if encodeWrr != nil {
		logger.Error("commonHandler=error=", encodeWrr)
	}
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
