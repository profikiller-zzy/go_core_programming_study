package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "go_core_programming/rpcStudy/stream_rpc_test/proto"

	"google.golang.org/grpc"
)

const Port = ":50052"

type server struct {
	pb.UnimplementedGreeterServer
}

// GetStream 实现服务端流模式
func (s *server) GetStream(req *pb.StreamRequest, serverStr pb.Greeter_GetStreamServer) error {
	// 示例：持续发送数据，直到上下文被取消或达到条件
	for i := 1; i <= 10; i++ {
		// 检查客户端是否已断开连接
		if errors.Is(serverStr.Context().Err(), context.Canceled) {
			return serverStr.Context().Err()
		}

		// 构造响应数据
		resp := &pb.StreamResponse{
			Data: fmt.Sprintf("Stream Message #%d", i),
		}

		// 发送数据到客户端
		if err := serverStr.Send(resp); err != nil {
			// 处理发送错误（如客户端断开）
			return err
		}

		// 模拟延迟（例如每秒发送一次）
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *server) PutStream(clientStr grpc.ClientStreamingServer[pb.StreamRequest, pb.StreamResponse]) error {
	// 不断调用Recv()方法接收客户端发送的数据流
	for {
		req, err := clientStr.Recv()
		if err != nil { // 客户端数据流发送成功或者发生错误
			log.Printf("%v\n", err)
			break
		}
		fmt.Println("接收到的数据: ", req.Data)
	}
	clientStr.SendAndClose(&pb.StreamResponse{Data: "PutStream Response"})
	return nil
}

func (s *server) AllStream(allStream pb.Greeter_AllStreamServer) error {
	// 这里就不能像之前的 GetStream 一样，直接 for 循环发送数据了
	// 我们定义两个协程，一个协程用于发数据一个用于接收数据
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			// 不断的接收客户端发送的数据
			data, err := allStream.Recv()
			if err != nil {
				log.Printf("接收数据失败: %v\n", err)
				break
			}
			fmt.Println("接收到的数据: ", data.Data)
		}
	}()

	go func() {
		defer wg.Done() // 协程结束时减少计数
		for i := 0; i <= 10; i++ {
			_ = allStream.Send(&pb.StreamResponse{Data: "你好我是服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait() // 等待两个协程结束
	return nil
}

func main() {
	//s := server{}
	//t := reflect.TypeOf(s)
	//fmt.Println("server 结构体的方法列表:")
	//for i := 0; i < t.NumMethod(); i++ {
	//	fmt.Println("-", t.Method(i).Name)
	//}
	//
	//fmt.Println("----------------------------------")
	//
	//t = reflect.TypeOf(pb.UnimplementedGreeterServer{})
	//fmt.Println("pb.UnimplementedGreeterServer{} 结构体的方法列表:")
	//for i := 0; i < t.NumMethod(); i++ {
	//	fmt.Println("-", t.Method(i).Name)
	//}

	// 1. 服务器首先监听端口
	listener, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}

	// 2. 创建 gRPC 服务器
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// 3. 服务端开始提供服务，等待客户端连接并调用
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
