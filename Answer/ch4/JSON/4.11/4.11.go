package _4_11

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type PosterResp struct {
	Poster string
	Title  string
}

type XkcdResp struct {
	Img   string
	Title string
	Link  string
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CategorizedCreatedAt(issues []*Issue) {
	now := time.Now().UTC()
	var OverYear, LessYear, LessMonth []*Issue

	for _, issue := range issues {
		issueTime := issue.CreatedAt.UTC()
		isOverOneYear := now.Year()-issueTime.Year() > 1
		isLessOneMonth := !isOverOneYear && (now.Month()-issueTime.Month() < 1)
		if isOverOneYear {
			OverYear = append(OverYear, issue)
		} else if !isOverOneYear {
			LessYear = append(LessYear, issue)
		} else if isLessOneMonth {
			LessMonth = append(LessMonth, issue)
		}
	}
	for _, issue := range OverYear {
		fmt.Printf("Over one Year: #%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
	for _, issue := range LessYear {
		fmt.Printf("Less one Year: #%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
	for _, issue := range LessMonth {
		fmt.Printf("Less one Month: #%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}

}

func parseBody(response *http.Response) ([]byte, error) {
	// 调用ReadAll函数将`io.ReadCloser`接口的数据流全部读入字节切片中
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func queryUrl(url string) ([]byte, error) {
	// 1. query url
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("query url error: %v", err)
		return nil, err
	}
	// 2. parse response body
	body, err := parseBody(resp)
	if err != nil {
		fmt.Printf("parse response error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

func Poster() {
	// 1. get movie name
	// you can apply free apiKey
	ImageUrl := "http://www.omdbapi.com/?apikey=[your key]&"
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
	movieBody, err := queryUrl(movieUrl)
	if err != nil {
		return
	}
	// 3. download image
	var jsonMovieBody PosterResp
	json.Unmarshal(movieBody, &jsonMovieBody)
	imgUrl := jsonMovieBody.Poster
	imgBody, err := queryUrl(imgUrl)
	if err != nil {
		return
	}
	// 4. write file
	var validSuffix = regexp.MustCompile(`\.(jpe?g|web|png|gif)$`)
	suffix := validSuffix.FindString(imgUrl)
	fileName := string(movieName) + suffix
	fileErr := ioutil.WriteFile(fileName, imgBody, 0644)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
}

func Xkcd() {
	xkcdUrl := "https://xkcd.com/"
	xkcdSuffix := "/info.0.json"
	f, err := os.OpenFile("storage.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var max int
	fmt.Println("------------------- how match you want -----------------------")
	fmt.Scanln(&max)

	for i := 1; i < max; i++ {
		url := xkcdUrl + fmt.Sprint(i) + xkcdSuffix
		fmt.Printf("query url: %v\n", url)
		body, err := queryUrl(url)
		if err != nil {
			fmt.Printf("query url error: %v\n", err)
			break
		}
		var parseBody XkcdResp
		if err := json.Unmarshal(body, &parseBody); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
			continue
		}
		row := parseBody.Img + parseBody.Title + parseBody.Link + "\n"
		if _, err := f.Write([]byte(row)); err != nil {
			fmt.Printf("write info error: %v\n", err)
		}
	}
	f.Close()

	rf, rErr := os.OpenFile("storage.txt", os.O_RDONLY, 0)
	if rErr != nil {
		log.Fatal(err)
	}

	var order int
	fmt.Println("------------------- want order -----------------------")
	fmt.Scanln(&order)

	reader := bufio.NewReader(rf)
	flag := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF: %#v\n", line)
				break
			}
		}
		if flag == order-1 {
			fmt.Printf("%v", line)
		}
		flag++
	}

	defer f.Close()
}

func queryIssue(address, method string) {
	// 1. generator template
	f, err := os.OpenFile("template.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	template := "title:\nbody:\nstate:\n"
	_, wErr := f.WriteString(template)
	if err != nil {
		fmt.Printf("generator template error: %v", wErr)
		return
	}
	f.Close()

	// 2. open template, waiting for user write
	// Attention: use case run on mac os
	cmd := exec.Command("open", "template.txt")
	if err := cmd.Run(); err != nil {
		os.Exit(1)
		log.Fatal(err)
		return
	}

	var moveOn string
	fmt.Println("-- waiting fo input title, body and state, type continue to move on --")
	fmt.Scanln(&moveOn)
	for moveOn != "continue" {
	}

	// 3. read template, get title body etc.
	f, ferr := os.OpenFile("template.txt", os.O_RDONLY, 0666)
	if ferr != nil {
		log.Fatal(ferr)
		return
	}
	reader := bufio.NewReader(f)
	lines := make(map[string]string)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		parseLine := string(line)
		if strings.Contains(parseLine, ":") {
			arr := strings.Split(parseLine, ":")
			if arr[1] != "" {
				lines[arr[0]] = arr[1]
			}
		}
	}
	defer f.Close()

	if len(lines) > 0 {
		jsonBytes, err := json.MarshalIndent(lines, "", " ")
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, address, strings.NewReader(string(jsonBytes)))
		req.Header.Add("Content-Type", "application/json")
		// https://docs.github.com/en/rest/overview/authenticating-to-the-rest-api?apiVersion=2022-11-28
		req.Header.Add("Authorization", "Bearer [your token]")
		resp, err := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%v \n %v \n %v", address, string(body), strings.NewReader(string(jsonBytes)))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Success: %v", string(body))
		}
		defer resp.Body.Close()
	} else {
		fmt.Println("nothing input, quit")
	}
}

func IssueOperation() {
	var owner string
	var repo string
	var issueNumber string
	var operationType int
	const OperationURL = "https://api.github.com/repos/"
	flag.IntVar(&operationType, "operation", 0, "Specify operation type: 1 -- create; 2 -- get; 3 -- update; 4 -- delete")
	flag.StringVar(&owner, "owner", "", "The account owner of the repository. The name is not case sensitive.")
	flag.StringVar(&repo, "repo", "", "The name of the repository without the .git extension. The name is not case sensitive.")
	flag.StringVar(&issueNumber, "issue_n", "", "The number that identifies the issue.")
	flag.Parse()

	if owner != "" && repo != "" && operationType != 0 {
		var queryAddress string

		switch operationType {
		case 1:
			// command: go run hello.go -operation=1 -owner=[your owner] -repo=[your repo]
			queryAddress = OperationURL + owner + "/" + repo + "/issues"
			queryIssue(queryAddress, "POST")
		case 2:
			// command go run hello.go -operation=2 -owner=golang -repo=go -issue_n=5680
			queryAddress = OperationURL + owner + "/" + repo + "/issues/" + issueNumber
			body, err := queryUrl(queryAddress)
			if err != nil {
				fmt.Printf("get issue error: %v\n", err)
				return
			}
			var result Issue
			parseErr := json.Unmarshal(body, &result)
			if parseErr != nil {
				fmt.Printf("parse body error: %v", parseErr)
				return
			}
			fmt.Println(result)
		case 3:
			// command: go run hello.go -operation=3 -owner=[your owner] -repo=[your repo] -issue_n=1
			queryAddress = OperationURL + owner + "/" + repo + "/issues/" + issueNumber
			queryIssue(queryAddress, "PATCH")
		case 4:
			// command: go run hello.go -operation=4 -owner=[your owner] -repo=[your repo] -issue_n=1
			// set state -> 4
			queryAddress = OperationURL + owner + "/" + repo + "/issues/" + issueNumber
			queryIssue(queryAddress, "PATCH")
		}
	} else {
		fmt.Println("application need correct input")
	}
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	// Remember to comment other code
	fmt.Println("---------------- 4.10 Categorized by time start -----------------")
	CategorizedCreatedAt(result.Items)
	fmt.Println("---------------- 4.10 Categorized by time end -------------------")

	// Remember to comment other code
	fmt.Println("---------------- 4.11 Issue start -----------------")
	IssueOperation()
	fmt.Println("---------------- 4.11 Issue end -------------------")

	// Remember to comment other code
	fmt.Println("---------------- 4.12 xkcd start -----------------")
	Xkcd()
	fmt.Println("---------------- 4.12 xkcd end -------------------")

	// Remember to comment other code
	fmt.Println("---------------- 4.13 Poster start -----------------")
	Poster()
	fmt.Println("---------------- 4.13 Poster end -------------------")
}
