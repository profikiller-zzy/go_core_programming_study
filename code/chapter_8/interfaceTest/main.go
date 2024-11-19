package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.WriteCloser
	w = os.Stdout
	buf := make([]byte, 16)
	buf = []byte("I love you")
	n, err := w.Write(buf)
	fmt.Println(n, err)
	w.Close()

	var w1 io.Writer
	w1 = os.Stdout
	w1.Write([]byte("I love you"))

	var any io.Writer
	any = os.Stdout
	any.Write(buf)
}
