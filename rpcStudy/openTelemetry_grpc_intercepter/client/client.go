package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "go_core_programming/rpcStudy/metadata_test/proto"
)

func returnOpenTelemetryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("Intercepting request to method:", method)
		return nil
	}
}

func main() {
	// 1. 初始化拦截器
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(returnOpenTelemetryClientInterceptor()))

	// 2. 创建连接，设置拦截器
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println("Server Response:", resp.Message)
}
