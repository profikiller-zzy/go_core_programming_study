package main

import "net"

// 客户端 - 快速发送多个包
func client() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	// 快速连续发送三个包
	conn.Write([]byte("Hello"))
	conn.Write([]byte("World"))
	conn.Write([]byte("Goodbye"))
}

func main() {
	client()
}
