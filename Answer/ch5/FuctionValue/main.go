package main

import (
	"GoStudy/Answer/ch5/FuctionValue/5.8"
	"GoStudy/Answer/ch5/FuctionValue/5.9"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://gopl-zh.github.io/index.html"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("get url : %s error :%s\n", url, err)
		resp.Body.Close()
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad response code : %s\n", resp.Status)
		resp.Body.Close()
		os.Exit(1)
	}
	// 2. parse response body
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("parse response error: %v", err)
		resp.Body.Close()
		os.Exit(1)
	}
	resp.Body.Close()
	//_5_7.ForEachNode(doc, _5_7.StartElement, _5_7.EndElement)
	resNode := _5_8.ElementByID(doc, _5_8.Pre, "img")
	fmt.Println(resNode)

	s := "foofoosfooThis is a string containing foo."
	fmt.Println(_5_9.Expand(s, _5_9.F))
}
