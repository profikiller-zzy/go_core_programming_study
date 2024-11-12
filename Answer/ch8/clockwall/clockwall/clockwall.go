package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s City=Host:Port...\n", os.Args[0])
		os.Exit(1)
	}

	locations := make(map[string]string) // 地区和socket之间的对应关系
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invaild argument: %v\n", arg)
			os.Exit(1)
		}
		locations[parts[0]] = parts[1]
	}

	for city, address := range locations {
		// 使用匿名函数
		go func(city, addr string) {
			var conn net.Conn
			var err error
			for attempt := 1; attempt <= 10; attempt++ {
				conn, err = net.Dial("tcp", addr)
				if err == nil {
					break
				}
				// 没有正确与该服务器建立连接，大概率是没有开启相应端口的clock2服务，所以允许最多与每个服务器尝试重连十次
				fmt.Fprintf(os.Stderr, "Failed to connect to %s (%s): %v. Retrying %d/10\n", city, addr, err, attempt)
				time.Sleep(1 * time.Second)
			}
			defer func() { // 延迟关闭
				if conn != nil {
					conn.Close()
				}
			}()
			if conn == nil { // 尝试十次后都没有连接成功，直接返回
				return
			}
			// 每隔一秒将服务器中读到的信息输出到控制台
			buf := make([]byte, 1024) // 缓存大小
			for {
				n, err := conn.Read(buf)
				if err != nil {
					if err != io.EOF {
						log.Printf("Failed to read time from %s: %v\n", city, err)
					}
					return
				}
				fmt.Fprintf(os.Stdout, "City: %s; Time: %s", city, string(buf[:n]))
				time.Sleep(1 * time.Second)
			}
		}(city, address)
	}

	// 创建一个空的select来永远阻塞main goroutine的退出
	select {}
}
