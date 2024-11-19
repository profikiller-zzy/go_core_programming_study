package webCal

import (
	"fmt"
	"go_core_programming/Answer/ch7/7_9_Expr/eval"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

// 练习 7.16： 编写一个基于web的计算器程序。

type WebCalculator struct {
	expr eval.Expr
	env  eval.Env
}

func (wc *WebCalculator) ListenAndServe() {
	// 设置 HTTP 路由
	http.HandleFunc("/setEnv", wc.setEnv)             // 设置环境变量
	http.HandleFunc("/parseExpr", wc.parseExprString) // 解析表达式
	http.HandleFunc("/calculate", wc.returnVal)       // 计算并返回结果

	// 启动服务器
	fmt.Println("服务器启动，访问 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (wc *WebCalculator) parseExprString(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			// 处理解析期间的 panic 错误
			http.Error(w, fmt.Sprintf("An error occurred in parsing the expression: %v\n", err), http.StatusInternalServerError)
		}
	}()

	// 获取表达式字符串
	exprString := r.URL.Query().Get("expr")

	// 调用 Parse 解析表达式并处理可能的错误
	parsedExpr, err := eval.Parse(exprString)
	if err != nil {
		// 如果解析错误，返回错误信息给用户
		http.Error(w, fmt.Sprintf("Failed to parse expression: %v\n", err), http.StatusBadRequest)
		return
	}

	// 如果解析成功，将解析结果存储在 webCalculator 实例中
	wc.expr = parsedExpr
}

func (wc *WebCalculator) setEnv(w http.ResponseWriter, r *http.Request) {
	// 创建一个新的环境
	env := make(eval.Env)

	// 从 URL 查询参数中解析变量
	query := r.URL.Query()
	for key, values := range query {
		if len(values) > 0 {
			// 将字符串值转换为浮点数
			value, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid value for %s: %v", key, err), http.StatusBadRequest)
				return
			}
			// 设置环境变量
			env[eval.Var(key)] = value
		}
	}

	// 设置计算器的环境
	wc.env = env

	// 响应成功消息
	fmt.Fprintf(w, "Environment variables set successfully: %v", env)
}

func (wc *WebCalculator) returnVal(w http.ResponseWriter, r *http.Request) {
	if wc.expr == nil || (reflect.ValueOf(wc.expr).Kind() == reflect.Ptr && reflect.ValueOf(wc.expr).IsNil()) {
		// expr 的动态值为 nil
		fmt.Fprintf(w, "Expr is empty, please input Expr first")
	}
	// 假设 Eval 返回 float64 并需要一个空的环境
	result := wc.expr.Eval(wc.env) // 使用空的 Env
	fmt.Fprintf(w, "Value: %f", result)
}
