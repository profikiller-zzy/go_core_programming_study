package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，
// 从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。
// 如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，
// 这些实例监听不同的端口，假装自己在不同的时区。像下面这样：
//
//
// $ TZ=US/Eastern    ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 添加flag以支持读入端口
	portFlag := flag.String("port", "8000", "the port to listen on")
	flag.Parse()

	// 设置时区，从环境变量中获取时区
	tz := os.Getenv("TZ")
	if tz == "" {
		log.Fatal("TZ environment variable is not set")
	}

	// 加载指定时区
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("Invalid timezone specified: %v", err)
	}

	address := fmt.Sprintf("localhost:%s", *portFlag)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn, location)
	}
}
