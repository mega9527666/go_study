package grpc_server

import (
	"context"
	"fmt"
	"mega/engine/logger"
	"mega/grpc/jhao_grpc_proto"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	jhao_grpc_proto.UnimplementedGrpcServiceServer
}

func StartGrpcServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("StartGrpcServer error=", port, err)
		return err
	}

	grpcServer := grpc.NewServer()

	jhao_grpc_proto.RegisterGrpcServiceServer(grpcServer, &GrpcServer{})

	return grpcServer.Serve(lis)
}

func (s *GrpcServer) SayHello(
	ctx context.Context,
	req *jhao_grpc_proto.HelloRequest,
) (*jhao_grpc_proto.HelloResponse, error) {

	return &jhao_grpc_proto.HelloResponse{
		Message: "hello " + req.Name,
	}, nil
}
