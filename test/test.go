package main

import (
	"mega/common/config"
	"mega/engine/logger"
	"mega/grpc/grpc_client_manager"
	"time"
)

// "github.com/gorilla/websocket"

func main() {

	var http_port int = 9800
	config.Env = config.EnvDev
	config.ServerType = config.ServerType_List.Hall_server

	//http client请求
	// dataObj := map[string]any{
	// 	"v": "1.0.2",
	// 	"t": 123456789,
	// }
	// dataStr := json_util.Stringify(dataObj)
	// realParam := map[string]string{
	// 	"k":    md5_helper.GetMd5_encrypt(dataStr),
	// 	"data": dataStr,
	// }

	// client := http_client.NewHttpClient(0)
	// // GET
	// resp, err := client.Post(
	// 	// "https://ipinfo.io/json",
	// 	"http://127.0.0.1:9090/init_client",
	// 	realParam,
	// 	map[string]string{
	// 		"Content-Type": "application/x-www-form-urlencoded",
	// 	},
	// )

	// if err != nil {
	// 	logger.Warn("请求错误=", err)
	// } else {
	// 	// logger.Log("请求成功=", resp)
	// 	var dataObj map[string]interface{}
	// 	if err := json.Unmarshal([]byte(resp), &dataObj); err != nil {

	// 	} else {
	// 		logger.Log("请求成功=json--", dataObj)
	// 	}
	// }

	//读取yaml配置文件
	config.InitConfig(http_port)

	grpc_client_manager.InitGrpcClient()
	time.AfterFunc(3*time.Second, func() {
		logger.Log("延迟执行:")
		grpc_client_manager.Login()
	})

	// 关键：等待足够时间让定时器执行
	time.Sleep(5 * time.Second) // 等待 4 秒，确保 3 秒的定时器能执行

}
