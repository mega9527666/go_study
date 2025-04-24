package webreqhandler

import (
	"encoding/json"
	"io"
	"mega/engine/logger"
	"net/http"
	"os"
	"strconv"
)

type httpHandleFunc func(w http.ResponseWriter, r *http.Request)

var routesMap = map[string]httpHandleFunc{
	"/mega":       dispatcher(megaHandler),
	"/login_jhao": dispatcher(login_jhao_handler),
}

func init() {
	// logger.Log("webreqhandler.init")
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
func dispatcher(next httpHandleFunc) httpHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 允许所有来源，生产环境可改为指定域名
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头，这里加上你需要的自定义头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,X-token")
		// 如果你希望前端能够读取到某些响应头，则用 Expose-Headers
		w.Header().Set("Access-Control-Expose-Headers", "X-My-Custom-Header")

		// go indexHandler(w, r, next)
		var contentType string = r.Header.Get("content-type")
		if len(contentType) == 0 {
			r.Header.Set("content-type", "application/json")
		}
		logger.Log("dispatcher=param=Content-Type", r.Header.Get("content-type"))
		logger.Log("dispatcher=param=X-token", r.Header.Get("X-token"))
		logger.Log("dispatcher=param=x-token", r.Header.Get("x-token"))
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

// 定义示例结构体
// 根据你的实际需求调整字段类型和标签
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func login_jhao_handler(w http.ResponseWriter, r *http.Request) {
	logger.Log("login_jhao_handler==", r.RequestURI, r.Host, r.RemoteAddr)
	logger.Log("login_jhao_handler=param=Body", r.Body)
	// logger.Log("login_jhao_handler=param=Header", r.Header)
	// logger.Log("login_jhao_handler=param=Content-Type", r.Header.Get("Content-Type"))
	logger.Log("login_jhao_handler=param=Body", r.Body)
	var queryMap = r.URL.Query()
	logger.Log("login_jhao_handler=param=queryMap", queryMap)
	logger.Log("login_jhao_handler=param=gameId", queryMap.Get("gameId"))
	// 设置响应的 Content-Type 为 text/plain
	// w.Header().Set("Content-Type", "text/plain")
	// w.Header().Set("Content-Type", "application/json")

	// 示例：将结构体转换为 JSON（序列化）
	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		logger.Warn("JSON 序列化失败: %v", err)
	}
	logger.Log("序列化后的 JSON 字符串: %s\n", string(jsonData))

	// 示例：将 JSON 字节切片转换为结构体（反序列化）
	data := []byte(`{"name":"Bob","age":25,"email":"bob@example.com"}`)
	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		logger.Warn("JSON 反序列化失败: %v", err)
	}
	logger.Log("反序列化后的结构体: %+v\n", u)

	// 1. 延迟关闭 Body
	defer r.Body.Close()

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "读取请求体失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// // 解析 JSON
	// var reqBody RequestBody
	// if err := json.Unmarshal(body, &reqBody); err != nil {
	// 	http.Error(w, "JSON 解析失败: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// 3. 转为字符串（如果需要）
	// bodyStr := string(data)
	logger.Log("Request body:", body)
	logger.Log("Request string body:", string(body))

	// 4. 响应
	// w.WriteHeader(http.StatusOK)

	// 创建一个响应对象
	response := Response{Message: "Hello, JSON abcd!"}
	// 向客户端写入响应内容
	w.WriteHeader(http.StatusOK)
	// 将结构体编码为 JSON 并写入响应体
	encodeWrr := json.NewEncoder(w).Encode(response)
	// _, err := w.Write([]byte("Hello, abcd!"))
	if encodeWrr != nil {
		logger.Error("login_jhao_handler=error=", encodeWrr)
	}
}
