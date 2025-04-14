package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "go_core_programming/rpcStudy/metadata_test/proto"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %s", req.Name)
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

// returnIntercept 拦截器函数，（单次请求拦截器） ，
func returnIntercept() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 1. 首先提取metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("metadata not found")
			return &pb.HelloReply{Message: "metadata not found"}, nil
		}
		for key, value := range md {
			fmt.Println(key, value)
		}
		// 如果不想直接返回，可能还有其他逻辑的话，就先接收handler的返回值
		return handler(ctx, req)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 把metadata的验证写入到拦截器中，首先创建拦截器
	myInterceptor := returnIntercept()

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(myInterceptor))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
