package main

import (
	"context"
	"google.golang.org/grpc"
	pb "rpcStudy/helloworld_rpc_test/proto"
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
	client.SayHello(context.Background(), &pb.User{Name: "test"})
}
