package grpc_client_manager

import (
	"context"
	"mega/common/config"
	"mega/engine/logger"
	"mega/grpc/grpc_client"
	"sync"
)

var (
	Web_server_GrpcClient_Map     map[string]*grpc_client.GrpcClient
	Account_server_GrpcClient_Map map[string]*grpc_client.GrpcClient
	Hall_server_GrpClient_Map     map[string]*grpc_client.GrpcClient

	once sync.Once
)

func InitGrpcClient() {
	once.Do(func() {
		initAll()
	})
}

func initAll() {
	// 默认初始化为空 map，防止 nil panic
	Web_server_GrpcClient_Map = make(map[string]*grpc_client.GrpcClient)
	Account_server_GrpcClient_Map = make(map[string]*grpc_client.GrpcClient)
	Hall_server_GrpClient_Map = make(map[string]*grpc_client.GrpcClient)

	if config.ServerType != config.ServerType_List.Web_server {
		Web_server_GrpcClient_Map = initClientMap(
			config.Global_Config.WebServer,
		)
	}

	if config.ServerType != config.ServerType_List.Account_server {
		Account_server_GrpcClient_Map = initClientMap(
			config.Global_Config.AccountServer,
		)
	}

	if config.ServerType != config.ServerType_List.Hall_server {
		Hall_server_GrpClient_Map = initClientMap(
			config.Global_Config.HallServer,
		)
	}

	logger.Log("Web_server_GrpcClient_Map===", Web_server_GrpcClient_Map)
	logger.Log("Account_server_GrpcClient_Map=====", Account_server_GrpcClient_Map)
	logger.Log("Hall_server_GrpClient_Map====", Hall_server_GrpClient_Map)
}

func initClientMap(
	servers []config.ServerItem,
) map[string]*grpc_client.GrpcClient {

	m := make(map[string]*grpc_client.GrpcClient, len(servers))
	for _, v := range servers {
		c := grpc_client.NewGrpcClient(v)
		m[c.Addr] = c
	}
	return m
}

func Login() {
	for _, v := range Web_server_GrpcClient_Map {
		resp, err := v.SayHello(context.Background(), "from hall")
		if err != nil {
			continue
		}
		logger.Log("grpc Login=====", resp)
		break // 👉 只要一个成功就行
	}
}
