package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "rpcStudy/grpc_error_test/proto"

	"google.golang.org/grpc"
)

func main() {
	// 1. 创建连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 设置超时时间 10s
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
	if err != nil { // 处理错误
		st, ok := status.FromError(err)
		if !ok {
			// 这个error不是一个gRPC错误，可能是自定义类型的错误
			panic("不是grpc错误，解析错误失败")
		}
		fmt.Println("Code:", st.Code(), ";Message:", st.Message())
		return
	}
	fmt.Println("Server Response:", resp.Message)
}
