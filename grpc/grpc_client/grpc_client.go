package grpc_client

import (
	"context"
	"mega/common/config"
	"mega/engine/logger"
	"mega/grpc/jhao_grpc_proto"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	ServerItem config.ServerItem
	Addr       string
	once       sync.Once
	conn       *grpc.ClientConn
	client     jhao_grpc_proto.GrpcServiceClient
	timeout    time.Duration
}

// 构造函数
func NewGrpcClient(serverItem config.ServerItem) *GrpcClient {
	var addr string = serverItem.InternalIP + ":" + strconv.Itoa(serverItem.GrpcPort)
	return &GrpcClient{
		ServerItem: serverItem,
		Addr:       addr,
		timeout:    10 * time.Second,
	}
}

func (g *GrpcClient) InitOnce() error {
	var initErr error
	g.once.Do(func() {
		initErr = g.init()
	})
	return initErr

}

// 真正初始化（只会执行一次）
func (g *GrpcClient) init() error {
	conn, err := grpc.NewClient(
		g.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("GrpcClient init error", g.Addr, err)
		return err
	}
	g.conn = conn
	g.client = jhao_grpc_proto.NewGrpcServiceClient(conn)
	return nil
}

// 对外方法（你真正调用的）
func (g *GrpcClient) SayHello(ctx context.Context, name string) (string, error) {
	if err := g.InitOnce(); err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(ctx, g.timeout)
	defer cancel()

	resp, err := g.client.SayHello(ctx, &jhao_grpc_proto.HelloRequest{
		Name: name,
	})

	if err != nil {
		logger.Log("SayHello error", err)
		return "", err
	}
	logger.Log("SayHello suc=", resp.Message)
	return resp.Message, nil
}
