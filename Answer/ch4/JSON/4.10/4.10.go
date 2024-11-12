package _4_10

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
// issues程序位于/ch4/issues

// IssuesURL 请求的地址
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult 响应结构体
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
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues 获取响应数据
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 判断状态码
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

// Process 处理数据
func Process(data *IssuesSearchResult) {
	hash := make(map[string][]*Issue, 3)
	hash["不到一个月"] = make([]*Issue, 0)
	hash["不到一年"] = make([]*Issue, 0)
	hash["超过一年"] = make([]*Issue, 0)

	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)
	for _, value := range data.Items {
		if value.CreatedAt.After(oneMonthAgo) { // 假如item创建的时间点晚于一个月前，则是不到一个月
			hash["不到一个月"] = append(hash["不到一个月"], value)
		} else if value.CreatedAt.After(oneYearAgo) {
			hash["不到一年"] = append(hash["不到一年"], value)
		} else {
			hash["超过一年"] = append(hash["超过一年"], value)
		}
	}
	fmt.Printf("%d issues:\n", data.TotalCount)
	for key, value := range hash {
		fmt.Printf("%s:\n", key)
		for _, item := range value {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
