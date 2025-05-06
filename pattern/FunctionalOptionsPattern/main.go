package main

import (
	"fmt"
	"sync"

	"go_core_programming/pattern/FunctionalOptionsPattern/server"
)

var once sync.Once

func main() {
	once.Do(func() {})
	s := server.NewServer(server.WithHost("localhost"), server.WithPort(8888), server.WithTLS(false))
	fmt.Println(s)
}
