package main

import (
	"mega/common/config"
	"mega/engine/logger"
	"mega/engine/socket_common"
	"mega/engine/socket_connection"
	"mega/hall_server/hall_socket_msg_mgr"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http_port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Error("初始化端口失败:", os.Args, err)
		return
	}
	env, err := strconv.Atoi(os.Args[2])
	if err != nil {
		logger.Error("初始化环境失败:", os.Args, err)
		return
	}

	config.Env = env
	config.ServerType = config.ServerType_List.Hall_server
	//读取yaml配置文件
	config.InitConfig(http_port)
	logger.Info("Hall_server.main", os.Args, env, http_port)

	// http.HandleFunc("/ws", wsHandler)
	// http.HandleFunc("/ws", socket_common.WsHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		socket_common.WsHandler(w, r, onMessageHandler)
	})
	// log.Println("WebSocket 服务启动: ws://localhost:8080/ws")
	logger.Log("WebSocket 服务启动: ws://127.0.0.1:" + strconv.Itoa(http_port))
	http.ListenAndServe(":"+strconv.Itoa(http_port), nil)

}

func onMessageHandler(s *socket_connection.Socket_Connection,
	msgType int,
	data []byte) {
	logger.Log("onMessageHandler=====", msgType, string(data))
	hall_socket_msg_mgr.Login_resp(s)
}
