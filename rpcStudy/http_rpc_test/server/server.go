package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type helloService struct{}

func (s *helloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply指针的值来返回的
	*reply = "hello:" + request
	return nil
}

func main() {
	// 1. http调用需要一个URL，首先配置路由
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.ReadCloser
			io.Writer
		}{ReadCloser: r.Body, Writer: w}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})
	// 2. 注册处理逻辑 handler
	err := rpc.RegisterName("HelloService", &helloService{})
	if err != nil {
		log.Fatal("Error registering service:", err)
	}
	// 3. 启动服务
	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
