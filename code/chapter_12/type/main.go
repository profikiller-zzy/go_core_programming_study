package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w).String())
}
