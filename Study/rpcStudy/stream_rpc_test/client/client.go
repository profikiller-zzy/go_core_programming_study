package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "rpcStudy/stream_rpc_test/proto"
	"sync"
	"time"
)

func main() {
	// 1. 建立连接
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// 2.1 测试服务端流模式
	fmt.Println("----------服务端流模式测试----------")
	serverStream, err := client.GetStream(context.Background(), &pb.StreamRequest{Data: "stream rpc test"})
	if err != nil {
		panic(err)
	}
	// 这里服务器发送的数据流需要一直接收，接收一个就处理一个
	for {
		res, err := serverStream.Recv()
		if err != nil {
			fmt.Println("接收数据失败，err: ", err)
			break
		}
		fmt.Println("接收到的数据: ", res.Data)
	}

	// 2.2 测试客户端流模式
	fmt.Println("----------客户端流模式测试----------")
	clientStream, err := client.PutStream(context.Background())
	for i := 0; i <= 10; i++ {
		err := clientStream.Send(&pb.StreamRequest{Data: fmt.Sprintf("stream rpc test %d#", i)})
		if err != nil {
			fmt.Println("发送数据失败，err: ", err)
			break
		}
		time.Sleep(time.Second)
	}

	// 2.3 测试双向流模式
	// 首先获取双向流对象
	fmt.Println("----------双向流模式测试----------")
	allStream, err := client.AllStream(context.Background())
	if err != nil {
		fmt.Println("获取双向流对象失败，err: ", err)
		return
	}
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
			_ = allStream.Send(&pb.StreamRequest{Data: "你好我是客户端"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait() // 等待两个协程结束
}
