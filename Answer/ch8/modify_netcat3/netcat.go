// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

// !+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("not a TCP connection")
	}
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done\n")
		done <- struct{}{} // signal the main goroutine
	}()

	// 把标准输入的内容复制到连接中
	mustCopy(conn, os.Stdin)
	// 确保标准输入被关闭后只关闭写的部分
	tcpConn.CloseWrite()
	<-done // wait for background goroutine to finish
	// 完全关闭连接
	tcpConn.Close()
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
