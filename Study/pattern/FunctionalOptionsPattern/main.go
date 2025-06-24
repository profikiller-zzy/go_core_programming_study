package main

import (
	"fmt"
	"go_core_programming/Study/pattern/FunctionalOptionsPattern/server"
	"sync"
)

var once sync.Once

func main() {
	once.Do(func() {})
	s := server.NewServer(server.WithHost("localhost"), server.WithPort(8888), server.WithTLS(false))
	fmt.Println(s)
}
