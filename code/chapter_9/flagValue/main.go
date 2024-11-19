package main

import (
	"flag"
	"fmt"
)

type flagValue struct {
	user string // -u
	pwd  string // -pwd
	host string // -h
	port int    // -p
}

func main() {
	var f flagValue
	// 注册flag
	flag.StringVar(&f.user, "u", "", "用户名，默认为空")
	flag.StringVar(&f.pwd, "pwd", "", "密码, 默认为空")
	flag.StringVar(&f.host, "h", "localhost", "主机名, 默认为localhost")
	flag.IntVar(&f.port, "port", 3306, "端口号, 默认为3306")

	// 所有变量注册flag完毕之后，调用该方法解析
	flag.Parse()
	// 输出结果
	fmt.Printf("user=%v pwd=%v\nhost=%v port=%v\n",
		f.user, f.pwd, f.host, f.port)
}
