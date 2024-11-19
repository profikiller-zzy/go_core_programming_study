package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "code/chapter_9/fileOperation/test.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error: %q\n", err)
	}
	fmt.Printf("%v\n", string(data))
}

func writeFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	n, err := file.Write([]byte(data))
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}
	fmt.Printf("%d bytes writed", n)
	return nil
}
