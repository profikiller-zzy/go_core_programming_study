package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"google.golang.org/grpc"

	proto2 "go_core_programming/rpcStudy/mcp_client/proto"
)

var ChatClient proto2.WeatherServiceClient

func InitSrvConn() {
	servHost := "localhost"
	servPort := 50051

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", servHost, servPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ChatClient = proto2.NewWeatherServiceClient(conn)
}

// putStream 向chatStream发送对话
func putStream(wg *sync.WaitGroup, signal chan int, allStream grpc.BidiStreamingClient[proto2.ChatRequest, proto2.ChatResponse]) {
	defer func(allStream grpc.BidiStreamingClient[proto2.ChatRequest, proto2.ChatResponse]) {
		err := allStream.CloseSend()
		if err != nil {
			log.Printf("关闭发送流失败，err: %v\n", err)
			panic(err)
		}
	}(allStream)
	defer wg.Done()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("你: ")
		// 用户输入
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "quit" {
			fmt.Println("再见")
			break
		} else {
			err := allStream.Send(&proto2.ChatRequest{
				Message: userInput,
			})
			if err != nil {
				log.Printf("发送数据失败，err: %v\n", err)
				break
			}
		}
		<-signal // 用户输入完成之后卡在这里，等到模型回复之后继续往下运行
	}
}

func getStream(wg *sync.WaitGroup, signal chan int, allStream grpc.BidiStreamingClient[proto2.ChatRequest, proto2.ChatResponse]) {
	defer wg.Done()
	for {
		data, err := allStream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				fmt.Printf("接收消息流错误，err: %v", err)
				break
			}
		}
		fmt.Printf("🤖 OpenAI: %s\n", data.Response)
		signal <- 1
	}
}

func main() {
	InitSrvConn()
	stream, err := ChatClient.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("获取双向流对象失败: %v\n", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	sg := make(chan int, 1)

	go putStream(&wg, sg, stream)
	go getStream(&wg, sg, stream)
	wg.Wait()
}
