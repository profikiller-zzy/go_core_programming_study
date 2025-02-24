package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn) // 从服务器读取数据
	serverWriter := bufio.NewWriter(conn) // 向服务器发送数据

	for {
		// 读取用户输入的命令
		fmt.Print("ftp> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading command:", err)
			continue
		}
		command = strings.TrimSpace(command)
		cmd := strings.Split(command, " ")[0]
		if command == "" {
			continue
		}

		// 向服务器发送命令
		serverWriter.WriteString(command + "\n")
		serverWriter.Flush()

		switch cmd {
		case "close":
			fmt.Println("Closing connection")
			return
		case "get":
			// 接收文件数据
			filename := strings.TrimSpace(strings.TrimPrefix(command, "get"))
			if filename == "" {
				fmt.Println("Usage: get <filename>")
				continue
			}

			// 创建本地文件
			file, err := os.Create(filename)
			if err != nil {
				fmt.Printf("Error creating file: %v\n", err)
				continue
			}
			defer file.Close()

			// 从服务器读取文件数据并写入本地文件
			_, err = io.Copy(file, serverReader)
			if err != nil {
				fmt.Printf("Error receiving file data: %v\n", err)
				continue
			}
			fmt.Printf("File '%s' downloaded successfully.\n", filename)
		case "ls", "cd":
			for {
				line, err := serverReader.ReadString('\n')
				if err != nil {
					if err == io.EOF { // 读到结尾
						break
					}
					log.Println("Error reading from server:", err)
					break
				}
				if line == "END\n" {
					break
				}
				fmt.Print(line)
				// 如果行是提示符或表示响应结束，则退出循环
				if strings.HasPrefix(line, "ftp>") || strings.HasPrefix(line, "Closing connection") {
					break
				}
			}
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:2121")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	handleConn(conn)
}
