package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	pb "rpcStudy/metadata_test/proto"

	"google.golang.org/grpc"
)

// intercept 拦截器函数
func returnIntercept() func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 创建metadata
		md := metadata.New(map[string]string{
			"username": "bobby",
			"password": "123456",
		})
		// 将metadata附加到context中
		ctx = metadata.NewOutgoingContext(ctx, md)
		// 如果不想直接返回，可能还有其他逻辑的话，就先接收handler的返回值
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func main() {
	// 1. 初始化拦截器
	myIntercept := returnIntercept()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(myIntercept))
	opts = append(opts, grpc.WithInsecure())

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
