package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "rpcStudy/grpc_proto_test/proto"
	"time"
)

func main() {
	// 1. 建立连接
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewSayHelloClient(conn)
	// 传入一个空的对象，测试protobuf的默认值
	client.SayHello(context.Background(), &pb.HelloRequest{
		Name:      "test",
		CreatedAt: timestamppb.New(time.Now()),
	})
}
