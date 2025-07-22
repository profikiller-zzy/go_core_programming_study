package es_hotel

import (
	"fmt"
	"testing"
	"time"

	"go_core_programming/Study/es/es_hotel"
)

// TestSearchHotels 测试酒店搜索功能
func TestSearchHotels(t *testing.T) {
	// 初始化ES客户端
	typedESClient, err := es_hotel.InitTypesElastic()
	if err != nil {
		t.Fatalf("初始化ES客户端失败: %v", err)
	}

	// 等待ES连接稳定
	time.Sleep(2 * time.Second)

	// 测试用例1: 基础关键词搜索
	t.Run("基础关键词搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			Keyword:   "如家",
			Page:      1,
			Size:      10,
			Highlight: true,
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("搜索失败: %v", err)
			return
		}

		t.Logf("搜索结果: 总数=%d, 页码=%d, 每页=%d, 总页数=%d, 耗时=%dms",
			result.Total, result.Page, result.Size, result.Pages, result.Took)

		if len(result.Hotels) > 0 {
			t.Logf("第一个酒店: %+v", result.Hotels[0])
		}
	})

	// 测试用例2: 城市过滤搜索
	t.Run("城市过滤搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			City:      "北京",
			Page:      1,
			Size:      5,
			SortBy:    "score",
			SortOrder: "desc",
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("城市过滤搜索失败: %v", err)
			return
		}

		t.Logf("北京酒店搜索结果: 总数=%d", result.Total)

		// 验证所有结果都是北京的酒店
		for _, hotel := range result.Hotels {
			if hotel.City != "北京" {
				t.Errorf("期望城市为北京，实际为: %s", hotel.City)
			}
		}
	})

	// 测试用例3: 价格范围搜索
	t.Run("价格范围搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			MinPrice:  200,
			MaxPrice:  500,
			Page:      1,
			Size:      10,
			SortBy:    "price",
			SortOrder: "asc",
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("价格范围搜索失败: %v", err)
			return
		}

		t.Logf("价格范围搜索结果: 总数=%d", result.Total)

		fmt.Println(result)
		// 验证价格范围
		for _, hotel := range result.Hotels {
			if hotel.Price < 200 || hotel.Price > 500 {
				t.Errorf("酒店价格 %.2f 不在范围 [200, 500] 内", hotel.Price)
			}
		}
	})

	// 测试用例4: 品牌和星级组合搜索
	t.Run("品牌和星级组合搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			Brand:    "如家",
			StarName: "二钻",
			Page:     1,
			Size:     10,
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("品牌和星级组合搜索失败: %v", err)
			return
		}

		t.Logf("如家二钻酒店搜索结果: 总数=%d", result.Total)

		// 验证品牌和星级
		for _, hotel := range result.Hotels {
			if hotel.Brand != "如家" {
				t.Errorf("期望品牌为如家，实际为: %s", hotel.Brand)
			}
			if hotel.StarName != "二钻" {
				t.Errorf("期望星级为二钻，实际为: %s", hotel.StarName)
			}
		}
	})

	// 测试用例5: 地理位置搜索（如果有坐标数据）
	t.Run("地理位置搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			Latitude:  31.2,
			Longitude: 121.2,
			Distance:  "10km",
			Page:      1,
			Size:      5,
			SortBy:    "_geo_distance",
			SortOrder: "asc",
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("地理位置搜索失败: %v", err)
			return
		}

		t.Logf("地理位置搜索结果: 总数=%d", result.Total)
		for _, hotel := range result.Hotels {
			t.Logf("酒店: %v", hotel)
		}
	})

	// 测试用例6: 评分范围搜索
	t.Run("评分范围搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			MinScore:  4.0,
			MaxScore:  5.0,
			Page:      1,
			Size:      10,
			SortBy:    "score",
			SortOrder: "desc",
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("评分范围搜索失败: %v", err)
			return
		}

		t.Logf("高评分酒店搜索结果: 总数=%d", result.Total)

		// 验证评分范围
		for _, hotel := range result.Hotels {
			if hotel.Score < 4.0 || hotel.Score > 5.0 {
				t.Errorf("酒店评分 %.1f 不在范围 [4.0, 5.0] 内", hotel.Score)
			}
		}
	})

	// 测试用例7: 复合条件搜索
	t.Run("复合条件搜索", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			Keyword:   "酒店",
			City:      "上海",
			MinPrice:  300,
			MaxPrice:  800,
			MinScore:  3.5,
			Page:      1,
			Size:      10,
			Highlight: true,
			SortBy:    "score",
			SortOrder: "desc",
		}

		result, err := es_hotel.ExecuteSearch(req, typedESClient)
		if err != nil {
			t.Errorf("复合条件搜索失败: %v", err)
			return
		}

		t.Logf("复合条件搜索结果: 总数=%d", result.Total)

		// 验证复合条件
		for _, hotel := range result.Hotels {
			if hotel.City != "上海" {
				t.Errorf("期望城市为上海，实际为: %s", hotel.City)
			}
			if hotel.Price < 300 || hotel.Price > 800 {
				t.Errorf("酒店价格 %.2f 不在范围 [300, 800] 内", hotel.Price)
			}
			if hotel.Score < 3.5 {
				t.Errorf("酒店评分 %.1f 低于最低要求 3.5", hotel.Score)
			}
		}
	})

	// 测试用例8: 分页测试
	t.Run("分页测试", func(t *testing.T) {
		// 第一页
		req1 := es_hotel.SearchRequest{
			Page: 1,
			Size: 5,
		}

		result1, err := es_hotel.ExecuteSearch(req1, typedESClient)
		if err != nil {
			t.Errorf("第一页搜索失败: %v", err)
			return
		}

		// 第二页
		req2 := es_hotel.SearchRequest{
			Page: 2,
			Size: 5,
		}

		result2, err := es_hotel.ExecuteSearch(req2, typedESClient)
		if err != nil {
			t.Errorf("第二页搜索失败: %v", err)
			return
		}

		t.Logf("分页测试 - 第一页: %d条, 第二页: %d条", len(result1.Hotels), len(result2.Hotels))

		// 验证分页结果不同
		if len(result1.Hotels) > 0 && len(result2.Hotels) > 0 {
			if result1.Hotels[0].ID == result2.Hotels[0].ID {
				t.Error("分页结果重复，第一页和第二页返回了相同的酒店")
			}
		}
	})
}

// TestBuildSearchRequest 测试搜索请求构建
func TestBuildSearchRequest(t *testing.T) {
	t.Run("构建基础搜索请求", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			Keyword: "测试酒店",
			Page:    1,
			Size:    10,
		}

		searchReq, err := es_hotel.BuildSearchRequest(req)
		if err != nil {
			t.Errorf("构建搜索请求失败: %v", err)
			return
		}

		if searchReq == nil {
			t.Error("搜索请求为空")
			return
		}

		// 验证分页参数
		if *searchReq.From != 0 {
			t.Errorf("期望From为0，实际为: %d", *searchReq.From)
		}
		if *searchReq.Size != 10 {
			t.Errorf("期望Size为10，实际为: %d", *searchReq.Size)
		}

		t.Log("搜索请求构建成功")
	})

	t.Run("构建带过滤条件的搜索请求", func(t *testing.T) {
		req := es_hotel.SearchRequest{
			City:     "北京",
			Brand:    "如家",
			MinPrice: 200,
			MaxPrice: 500,
			Page:     2,
			Size:     20,
		}

		searchReq, err := es_hotel.BuildSearchRequest(req)
		if err != nil {
			t.Errorf("构建过滤搜索请求失败: %v", err)
			return
		}

		if searchReq == nil {
			t.Error("搜索请求为空")
			return
		}

		// 验证分页参数
		if *searchReq.From != 20 {
			t.Errorf("期望From为20，实际为: %d", *searchReq.From)
		}
		if *searchReq.Size != 20 {
			t.Errorf("期望Size为20，实际为: %d", *searchReq.Size)
		}

		t.Log("带过滤条件的搜索请求构建成功")
	})
}
