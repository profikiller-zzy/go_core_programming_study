package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	tick := time.NewTicker(10 * time.Second) // 创建一个定时器，每 10 秒检查一次
	inputChannel := make(chan string)

	go func() {
		defer close(inputChannel)
		for input.Scan() {
			inputChannel <- input.Text() // 将客户端传来的string传给channel
		}
		if err := input.Err(); err != nil {
			log.Print("Scanner error: ", err)
		}
	}()

	// 在循环中监听消息
	for {
		select {
		case <-tick.C: // 超时了，10秒没有消息
			fmt.Fprintln(c, "10秒没有发消息，超时了")
			return // 断开连接
		case <-time.After(10 * time.Second): // 继续等待输入，如果超过 10 秒没有输入，断开连接
			// 如果等待 10 秒都没有输入，直接断开连接
		case <-inputChannel: // 接收到客户端消息
			// 处理消息并进行回显
			echo(c, input.Text(), 1*time.Second)
			tick.Reset(10 * time.Second) // 重置定时器，客户端发来消息，重新计时
		}
	}

	if err := input.Err(); err != nil {
		log.Print("Scanner error: ", err)
	}
}
