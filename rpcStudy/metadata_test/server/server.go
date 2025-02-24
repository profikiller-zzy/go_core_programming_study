package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	pb "rpcStudy/metadata_test/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request from: %s", req.Name)
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 把metadata的验证写入到拦截器中，首先创建拦截器
	myInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
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

	opt := grpc.UnaryInterceptor(myInterceptor)
	grpcServer := grpc.NewServer(opt)
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
