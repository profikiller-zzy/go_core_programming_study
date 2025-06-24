package main

import (
	"context"
	"go_core_programming/Study/rpcStudy/grpc_retry/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + req.Name}, nil
}

// returnIntercept 单词请求拦截器
func returnIntercept() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		time.Sleep(time.Second * 2) // 睡两秒模拟网络拥塞
		return handler(ctx, req)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opt := grpc.UnaryInterceptor(returnIntercept())
	grpcServer := grpc.NewServer(opt)
	proto.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
