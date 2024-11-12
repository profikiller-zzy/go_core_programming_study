package _4_14

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"
)

// 练习 4.14： 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。

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

const IssuesURL = "https://api.github.com/search/issues"

// 模板
var temp = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func GitHubServer() {
	http.HandleFunc("/", Handle)
	// 监听端口
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	// 1.解析模板
	issuesList := template.Must(template.New("issuesList").Parse(temp))
	// 2.获取请求参数
	var issue = &IssuesSearchResult{}
	q := r.URL.Query().Get("key")
	if q == "" {
		fmt.Fprintln(w, "请输入要查询的内容")
		return
	}
	fmt.Println(q)
	fmt.Println(url.QueryEscape(IssuesURL + "?q=" + q))
	resp, err := http.Get(url.QueryEscape(IssuesURL + "?q=" + q))
	if err != nil {
		resp.Body.Close()
		fmt.Fprintln(w, err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintln(w, resp.StatusCode)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(issue); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer resp.Body.Close()
	issuesList.Execute(w, issue)
}
