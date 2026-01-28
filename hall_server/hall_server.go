package main

import (
	"mega/common/config"
	"mega/engine/logger"
	"mega/engine/socket_common"
	"mega/engine/socket_connection"
	"mega/engine/socket_worker"
	"mega/grpc/grpc_server"
	"mega/hall_server/hall_socket_msg_mgr"
	"net/http"
	"os"
	"runtime"
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
	socket_worker.InitWorkerPool(runtime.NumCPU() * 4) //创建	并发度 ≈ CPU × (1 + 等待时间 / 计算时间)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		socket_common.WsHandler(w, r, onMessageHandler)
	})
	//启动grpc服务
	go func() {
		err := grpc_server.StartGrpcServer(
			config.Now_ServerItem.GrpcPort,
		)
		if err != nil {
			logger.Warn("grpc start error:", err)
		}
	}()

	http.ListenAndServe(":"+strconv.Itoa(http_port), nil)

}

func onMessageHandler(s *socket_connection.Socket_Connection,
	msgType int,
	data []byte) {
	logger.Log("onMessageHandler=====", msgType, string(data))
	hall_socket_msg_mgr.Login_resp(s)
}
