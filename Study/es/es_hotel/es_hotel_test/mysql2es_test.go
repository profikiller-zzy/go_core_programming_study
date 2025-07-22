package es_hotel_test

import (
	"context"
	"fmt"
	"go_core_programming/Study/es/es_hotel"
	"log"
	"testing"
)

func mysql2es_test(t *testing.T) {
	// 初始化ES客户端
	if err := es_hotel.InitElastic(); err != nil {
		log.Fatalf("初始化ES失败: %v", err)
	}

	// 初始化MySQL连接
	db, err := es_hotel.InitMySQL()
	if err != nil {
		log.Fatalf("初始化MySQL失败: %v", err)
	}

	// 创建ES索引
	if err := es_hotel.CreateHotelIndex(); err != nil {
		log.Fatalf("创建ES索引失败: %v", err)
	}

	// 从MySQL获取数据
	ctx := context.Background()
	hotels, err := es_hotel.GetHotelsFromMySQL(ctx, db)
	if err != nil {
		log.Fatalf("获取MySQL数据失败: %v", err)
	}

	if len(hotels) == 0 {
		fmt.Println("没有找到酒店数据")
		return
	}

	// 批量导入到ES
	if err := es_hotel.BulkIndexHotels(hotels); err != nil {
		log.Fatalf("批量导入ES失败: %v", err)
	}

	fmt.Printf("\n数据迁移完成！共导入 %d 条酒店数据到ES索引 '%s'\n", len(hotels), es_hotel.HotelIndex)
}
