package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

const templ = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<table border="1" cellspacing="0">
		<tr style='text-align: left'>
			<th>item</th>
			<th>price</th>
		</tr>
		{{range $item, $price := .}}
		<tr>
			<td>{{$item}}</td>
			<td>{{$price}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>`

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// 注意 `http.HandlerFunc` 是类型转换
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/read", http.HandlerFunc(db.read))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars // 数据库类型

//func (db database) list(w http.ResponseWriter, req *http.Request) { // 处理函数，列出所有商品价格
//	for item, price := range db {
//		fmt.Fprintf(w, "%s: %s\n", item, price)
//	}
//}

// 练习 7.12： 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。html/template包（§4.6）可能会对你有帮助。
func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.New("list").Parse(templ)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "格式化模板失败, err : %q", err)
		return
	}
	if err := tmpl.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) { // 处理函数，列出单个商品价格
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// 练习 7.11： 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）
func (db database) create(w http.ResponseWriter, r *http.Request) {
	// 1.获取请求参数
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	// 2.检查商品是否存在
	if _, ok := db[item]; ok { // 商品存在
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item : %q 已经存在", item)
		return
	}
	// 商品存在
	// 3.创建商品记录
	priceF, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price : %q 格式错误!", price)
		return
	}
	db[item] = dollars(priceF)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "创建商品记录 item: %q, price: %s 成功!", item, dollars(priceF))
}

func (db database) read(w http.ResponseWriter, r *http.Request) {
	// 1.获取请求参数
	item := r.URL.Query().Get("item")
	// 2.检查商品是否存在
	if _, ok := db[item]; !ok { // 商品不存在
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item : %q 商品不存在", item)
		return
	}
	// 商品存在
	// 3.返回商品信息
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item: %q, price: %s", item, db[item])
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	// 1.获取请求参数
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	// 2.检查商品是否存在
	if _, ok := db[item]; !ok { // 商品不存在
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item : %q 商品不存在", item)
		return
	}
	// 商品存在
	// 3.更新商品价格
	priceF, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price : %q 格式错误!", price)
		return
	}
	db[item] = dollars(priceF)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "更新成功!")
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	// 1.获取请求参数
	item := r.URL.Query().Get("item")
	// 2.检查商品是否存在
	if _, ok := db[item]; !ok { // 商品不存在
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item : %q 商品不存在", item)
		return
	}
	// 商品存在
	// 3.删除商品信息
	delete(db, item)
	fmt.Fprintf(w, "item %q deleted!", item)
}
