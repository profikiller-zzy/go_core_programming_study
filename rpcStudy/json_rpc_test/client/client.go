package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 建立连接
	conn, _ := net.Dial("tcp", "localhost:1234")

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply) // hello:hello
}
