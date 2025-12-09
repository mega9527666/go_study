package webreqhandler

import (
	"encoding/json"
	"io"
	"mega/engine/http_common"
	"mega/engine/logger"
	"net/http"
)

var routesMap = map[string]http_common.HttpHandleFunc{
	// "/mega":        http_common.Dispatcher(megaHandler),
	// "/login_jhao":  http_common.Dispatcher(login_jhao_handler),
	"/init_client": http_common.Dispatcher(init_client_handler),
}

func init() {
	// logger.Log("webreqhandler.init")
}

func ListenAndServe(port int) {
	http_common.ListenAndServe(port, routesMap)
}

func megaHandler(w http.ResponseWriter, r *http.Request, ip string) {
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

func login_jhao_handler(w http.ResponseWriter, r *http.Request, ip string) {
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
		logger.Warn("JSON 序列化失败: ", err)
	}
	logger.Log("序列化后的 JSON 字符串: ", string(jsonData))

	// 示例：将 JSON 字节切片转换为结构体（反序列化）
	data := []byte(`{"name":"Bob","age":25,"email":"bob@example.com"}`)
	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		logger.Warn("JSON 反序列化失败: ", err)
	}
	logger.Log("反序列化后的结构体: ", u)

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

func init_client_handler(w http.ResponseWriter, r *http.Request, ip string, dataObj map[string]interface{}) {
	logger.Log("init_client_handler=param=Body", ip)
}
