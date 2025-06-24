package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "rpcStudy/helloworld_rpc_test/proto"
)

type helloService struct {
	pb.UnimplementedSayHelloServer
}

func (s *helloService) SayHello(ctx context.Context, req *pb.User) (*pb.User, error) {
	// 返回值是通过修改reply指针的值来返回的
	fmt.Println(req)
	return req, nil
}

func main() {
	// 1. 创建监听
	listener, _ := net.Listen("tcp", ":50052")
	// 2. 创建grpc服务器
	server := grpc.NewServer()
	pb.RegisterSayHelloServer(server, &helloService{})
	// 3. 启动服务
	server.Serve(listener)
}
