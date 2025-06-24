package client_proxy

import (
	"net/rpc"
	"rpcStudy/newhelloworld_rpc_test/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

// NewHelloServiceClient 构造函数，这样的写法在go中非常常见，因为在go语言中没有类、对象和初始化的方法
func NewHelloServiceClient(conn *rpc.Client) *HelloServiceStub {
	return &HelloServiceStub{Client: conn}
}

func (c HelloServiceStub) Hello(request string, reply *string) error {
	return c.Client.Call(handler.HelloServiceName+".Hello", request, reply)
}
