package _4_12

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。
//下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。

type XkcdResp struct {
	Img   string `json:"img"`
	Title string `json:"title"`
	Link  string `json:"link"`
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
		fmt.Printf("parse body error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	return bodyData, nil
}

func XKCD() {
	// xkcdURL 请求的xckdURL
	xkcdURL := "https://xkcd.com/"
	// xkcdSuffix 请求的后缀
	xkcdSuffix := "/info.0.json"
	// os.O_RDWR表示以读写模式打开文件 os.O_CREATE表示如果文件不存在则创建
	f, err := os.OpenFile("storage.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("输入你想爬取多少个漫画的JSON信息")
	var num int
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		fmt.Printf("The %dth item is processing, please wait..\n", i)
		url := xkcdURL + fmt.Sprintf("%d", i) + xkcdSuffix
		bodyData, err := urlQuery(url)
		if err != nil {
			fmt.Printf("the %dth item query url error: %v", i, err)
			continue
		}
		var parseBody XkcdResp
		if err := json.Unmarshal(bodyData, &parseBody); err != nil {
			log.Fatal(fmt.Sprintf("JSON unmarshal failed, err :%s", err))
			continue
		}
		row := parseBody.Img + " " + parseBody.Title + " " + parseBody.Link + "\n"
		if _, err := f.Write([]byte(row)); err != nil {
			log.Fatal(fmt.Sprintf("Write file failed, err :%s", err))
			continue
		}
	}
	f.Close()

	// 这里使用 0 表示不设置额外的权限
	rFile, rErr := os.OpenFile("storage.txt", os.O_RDONLY, 0)
	if rErr != nil {
		log.Fatal(fmt.Sprintf("Open file failed, err :%s", err))
	}

	fmt.Println("你想查看哪一条漫画的信息")
	var order int
	fmt.Scanln(&order)

	reader := bufio.NewReader(rFile)
	count := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF: %#v\n", line)
				break
			}
		}
		if count == order-1 {
			fmt.Println(line)
			break
		}
		count++
	}

	rFile.Close()
}
