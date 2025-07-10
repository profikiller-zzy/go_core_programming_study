package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
)

func main() {
	// 连接到 ShardingSphere-Proxy 的端口（不是直接连 mysql）
	dsn := "root:@tcp(127.0.0.1:3309)/sharding_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer db.Close()

	// 插入多条数据
	for i := 1; i <= 10; i++ {
		orderID := int64(10000 + i)
		userID := rand.Intn(10) // 随机用户
		status := "init"
		_, err := db.Exec("INSERT INTO t_order (order_id, user_id, status) VALUES (?, ?, ?)",
			orderID, userID, status)
		if err != nil {
			log.Printf("插入失败: %v", err)
		} else {
			fmt.Printf("插入成功: order_id=%d, user_id=%d\n", orderID, userID)
		}
	}

	fmt.Println("插入完成！请登录 mysql1 / mysql2 查看物理表数据。")
}
