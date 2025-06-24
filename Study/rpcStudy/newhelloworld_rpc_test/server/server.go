package main

import (
	"net"
	"net/rpc"
	"rpcStudy/newhelloworld_rpc_test/handler"
)

type helloService struct{}

func (s *helloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply指针的值来返回的
	*reply = "hello:" + request
	return nil
}

func main() {
	// 1. 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑 handler
	err := rpc.RegisterName(handler.HelloServiceName, &helloService{})
	if err != nil {
		return
	}
	// 3. 启动服务
	for {
		conn, _ := listener.Accept() // 当一个连接建立时，返回一个conn
		go rpc.ServeConn(conn)
	}
}
