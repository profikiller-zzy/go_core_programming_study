package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	nowTime := time.Now()
	fmt.Printf("now=%v now type=%T\n", nowTime, nowTime)

	fmt.Println(nowTime.Format("2006-01-02 15:04:05"))

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
}
