package main

import (
	"context"
	"fmt"
	"go_core_programming/Study/rpcStudy/grpc_retry/proto"
	"log"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	retryOptions := []retry.CallOption{
		retry.WithMax(3),                       // 重试的最大次数
		retry.WithPerRetryTimeout(time.Second), // 每次重试的超时时间
		retry.WithCodes(codes.Unavailable, codes.Unknown, codes.DeadlineExceeded), // 收到什么错误码时才重试
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// 添加重试拦截器，配置重试选项
	opts = append(opts, grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(retryOptions...)))

	conn, err := grpc.Dial(":50051", opts...)
	if err != nil {
		panic(err)
	}

	client := proto.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	fmt.Println("Server Response:", resp.Message)
}
