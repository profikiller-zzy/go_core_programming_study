package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: program <TimeZone=Address> [<TimeZone=Address>...]")
		os.Exit(1)
	}

	for _, str := range os.Args[1:] { // 参数格式：NewYork=localhost:8010
		parts := strings.Split(str, "=")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Invalid argument format: %v. Expected format: TimeZone=Address\n", str)
			continue
		}
		tz := parts[0] // 时区
		addr := parts[1]
		// 启动 Goroutine 处理每个时区
		go func(tz, addr string) {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Printf("Failed to connect to %s (%s): %v\n", tz, addr, err)
				return
			}
			defer conn.Close()

			fmt.Fprintf(os.Stdout, "Connected to %s (%s):\n", tz, addr)
			mustCopy(os.Stdout, conn)
		}(tz, addr)
	}

	// 防止主 Goroutine过早退出
	select {}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("Error during copy: %v\n", err)
	}
}
