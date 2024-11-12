package main

import (
	"GoStudy/Answer/ch5/MultipleReturnValue/5.5"
	"fmt"
	"os"
)

func main() {
	fmt.Println("----------------5.5 start----------------")
	//for _, url := range os.Args[1:] {
	//	words, images, err := _5_5.CountWordsAndImages(url)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "get %s words and images error: %s", url, err)
	//		continue
	//	}
	//	fmt.Printf("url: %s, words: %d, images: %d\n", url, words, images)
	//}
	url := "https://gopl-zh.github.io/index.html"
	words, images, err := _5_5.CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get %s words and images error: %s", url, err)
		os.Exit(1)
	}
	fmt.Printf("url: %s, words: %d, images: %d\n", url, words, images)
}
