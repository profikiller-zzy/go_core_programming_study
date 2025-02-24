package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// 练习 8.2： 实现一个并发FTP服务器。服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，
// ls来列出目录内文件，get和send来传输文件，close来关闭连接。你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

type FTPServer struct {
	rootDir string
}

func (s *FTPServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	cwd := s.rootDir

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading command:", err)
			}
			return
		}

		command = strings.TrimSpace(command)
		args := strings.Split(command, " ")
		cmd := args[0]

		switch cmd {
		case "cd":
			if len(args) < 2 {
				writer.WriteString("Usage: cd <directory>\n")
			} else {
				dir := args[1]
				newPath := filepath.Join(cwd, dir)
				fileInfo, err := os.Stat(newPath)
				if fileInfo.IsDir() && err == nil {
					cwd = newPath
					writer.WriteString("END\n") // 发送结束符
				} else {
					writer.WriteString(fmt.Sprintf("找不到路径“%s”，因为该路径不存在。\n", dir))
					writer.WriteString("END\n") // 发送结束符
				}
			}
		case "ls":
			files, err := os.ReadDir(cwd)
			if err != nil {
				writer.WriteString(fmt.Sprintf("Error listing directory: %v\n", err))
			} else {
				for _, file := range files {
					writer.WriteString(file.Name() + "\n")
				}
				writer.WriteString("END\n") // 发送结束符
			}
		case "get":
			if len(args) < 2 {
				writer.WriteString("Usage: get <filename>\n")
			} else {
				filename := args[1]
				filePath := filepath.Join(cwd, filename)
				file, err := os.Open(filePath)
				if err != nil {
					writer.WriteString(fmt.Sprintf("Error opening file: %v\n", err))
				} else {
					io.Copy(writer, file)
					file.Close()
				}
			}
		case "send":
			if len(args) < 2 {
				writer.WriteString("Usage: send <filename>\n")
			} else {
				filename := args[1]
				filePath := filepath.Join(cwd, filename)
				file, err := os.Create(filePath)
				if err != nil {
					writer.WriteString(fmt.Sprintf("Error creating file: %v\n", err))
				} else {
					io.Copy(file, reader)
					file.Close()
				}
			}
		case "close":
			writer.WriteString("Closing connection\n")
			return
		default:
			writer.WriteString("Unknown command\n")
		}
		writer.Flush()
	}
}

func main() {
	var rootDir string
	if len(os.Args) < 2 {
		log.Println("Usage: ftp_server <root-directory>")
		rootDir = "C:\\workDir"
	} else {
		rootDir = os.Args[1]
	}

	listener, err := net.Listen("tcp", "localhost:2121")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	log.Println("FTP server listening on port 2121")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		server := &FTPServer{rootDir: rootDir}
		go server.handleConnection(conn)
	}
}
