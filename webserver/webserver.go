package main

import (
	"encoding/json"
	"fmt"
	"mega/engine/logger"
	"mega/webserver/webreqhandler"
	"net/http"
	"os"
	"strconv"
	// "github.com/gorilla/websocket"
)

func main() {
	logger.Log("webserver.main", os.Args[1], fmt.Sprintf("%T", os.Args[1]))
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("初始化端口失败:", os.Args, err)
		return
	}
	logger.Log("webserver.main", os.Args[1], fmt.Sprintf("%T", os.Args[1]), port)
	webreqhandler.ListenAndServe(port)

	// http.HandleFunc("/", megaHandler)

}

func megaHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("megaHandler=Host=", r.URL.Query())
	w.WriteHeader(http.StatusOK)
	// w.Write("abcd")
	// 创建一个响应对象
	// response := Response{Message: "Hello, JSON World!"}
	json.NewEncoder(w).Encode("hello,abcd")
}
