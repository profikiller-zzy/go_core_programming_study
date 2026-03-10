package main

import (
	"fmt"
	"net"
)

// 服务端 - 演示粘包现象
func server() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go func(c net.Conn) {
			defer c.Close()
			buffer := make([]byte, 1024)

			for {
				n, err := c.Read(buffer)
				if err != nil {
					break
				}
				fmt.Printf("接收到数据: %s\n", string(buffer[:n]))
				// 可能输出："HelloWorldGoodbye" (三个包粘在一起)
			}
		}(conn)
	}
}

func main() {
	server()
	fmt.Println("服务器已启动，等待连接...")
	select {} // 阻塞主 goroutine
}
