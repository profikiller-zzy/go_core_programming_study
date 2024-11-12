package _4_13

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

// 练习 4.13： 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。
// 编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。

type PosterResp struct {
	Poster string `json:"poster"`
	Title  string `json:"title"`
}

func urlQuery(url string) ([]byte, error) {
	// 1, query url
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("query url error: %v", err)
		return nil, err
	}
	// 2. parse body
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("parse body error: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	return bodyData, nil
}

func Poster() {
	// 1. get movie name
	// you can apply free apiKey
	ImageUrl := "http://www.omdbapi.com/?apikey=faa22d39&"
	fmt.Println("---------------- Input move name ----------------------------")
	reader := bufio.NewReader(os.Stdin)
	movieName, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("cannot get movie name")
		return
	}
	// 2. query movie info
	escapeName := url.QueryEscape(string(movieName))
	movieUrl := ImageUrl + "&plot=full&t=" + escapeName
	movieBody, err := urlQuery(movieUrl)
	if err != nil {
		return
	}
	// 3. download image
	var jsonMovieBody PosterResp
	json.Unmarshal(movieBody, &jsonMovieBody)
	imgUrl := jsonMovieBody.Poster
	imgBody, err := urlQuery(imgUrl)
	if err != nil {
		return
	}
	// 4. write file
	var validSuffix = regexp.MustCompile(`\.(jpe?g|web|png|gif)$`)
	suffix := validSuffix.FindString(imgUrl)
	fileName := string(movieName) + suffix
	// `0644`:文件权限，代表该文件被设置为读写文件
	fileErr := ioutil.WriteFile(fileName, imgBody, 0644)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
}
