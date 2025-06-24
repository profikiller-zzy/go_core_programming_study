package main

import (
	"fmt"
	"net/rpc"
	"rpcStudy/newhelloworld_rpc_test/client/client_proxy"
)

func main() {
	// 1. 建立连接
	conn, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	// 2. 实例化一个client
	client := client_proxy.NewHelloServiceClient(conn)

	// 3. 调用服务 这里的调用方式和之前的不同，这里是通过client_proxy包中的NewHelloServiceClient方法来调用的
	var reply string
	err = client.Hello("hello", &reply)
	//err := conn.Call(handler.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply) // hello:hello
}
