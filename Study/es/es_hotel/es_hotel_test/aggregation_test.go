package es_hotel

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"go_core_programming/Study/es/es_hotel"
)

// TestAggregationQueries 测试各种聚合查询
func TestAggregationQueries(t *testing.T) {
	// 初始化ES客户端
	typedESClient, err := es_hotel.InitTypesElastic()
	if err != nil {
		t.Fatalf("初始化ES客户端失败: %v", err)
	}

	// 等待ES连接稳定
	time.Sleep(2 * time.Second)

	// 测试1：城市聚合查询
	t.Run("城市聚合查询", func(t *testing.T) {
		req := es_hotel.AggregationRequest{
			Keyword: "酒店",
		}

		searchReq, err := es_hotel.BuildCityAggregationQuery(req)
		if err != nil {
			t.Errorf("构建城市聚合查询失败: %v", err)
			return
		}

		result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
		if err != nil {
			t.Errorf("执行城市聚合查询失败: %v", err)
			return
		}

		t.Logf("城市聚合查询结果: 总数=%d, 耗时=%dms", result.Total, result.Took)

		// 解析城市统计结果
		cityStats, err := es_hotel.ParseCityAggregationResult(result)
		if err != nil {
			t.Errorf("解析城市聚合结果失败: %v", err)
			return
		}

		for _, stat := range cityStats {
			t.Logf("城市: %s, 酒店数量: %d, 平均价格: %.2f, 最高价格: %.2f, 最低价格: %.2f",
				stat.City, stat.Count, stat.AvgPrice, stat.MaxPrice, stat.MinPrice)
		}

		// 打印原始聚合结果
		aggJSON, _ := json.MarshalIndent(result.Aggregations, "", "  ")
		t.Logf("原始聚合结果:\n %s", string(aggJSON))
	})

	// 测试2：多维度聚合查询
	t.Run("多维度聚合查询", func(t *testing.T) {
		req := es_hotel.AggregationRequest{
			City: "北京",
		}

		searchReq, err := es_hotel.BuildMultiDimensionAggregationQuery(req)
		if err != nil {
			t.Errorf("构建多维度聚合查询失败: %v", err)
			return
		}

		result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
		if err != nil {
			t.Errorf("执行多维度聚合查询失败: %v", err)
			return
		}

		t.Logf("多维度聚合查询结果: 总数=%d, 耗时=%dms", result.Total, result.Took)

		// 解析品牌统计结果
		brandStats, err := es_hotel.ParseBrandAggregationResult(result)
		if err != nil {
			t.Errorf("解析品牌聚合结果失败: %v", err)
			return
		}

		for _, stat := range brandStats {
			t.Logf("品牌: %s, 酒店数量: %d, 平均评分: %.2f",
				stat.Brand, stat.Count, stat.AvgScore)
		}

		// 打印原始聚合结果
		aggJSON, _ := json.MarshalIndent(result.Aggregations, "", "  ")
		t.Logf("原始聚合结果: %s", string(aggJSON))
	})

	// 测试3：范围聚合查询
	t.Run("范围聚合查询", func(t *testing.T) {
		req := es_hotel.AggregationRequest{
			Keyword: "酒店",
		}

		searchReq, err := es_hotel.BuildRangeAggregationQuery(req)
		if err != nil {
			t.Errorf("构建范围聚合查询失败: %v", err)
			return
		}

		result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
		if err != nil {
			t.Errorf("执行范围聚合查询失败: %v", err)
			return
		}

		t.Logf("范围聚合查询结果: 总数=%d, 耗时=%dms", result.Total, result.Took)

		// 解析价格区间统计结果
		priceRangeStats, err := es_hotel.ParsePriceRangeAggregationResult(result)
		if err != nil {
			t.Errorf("解析价格区间聚合结果失败: %v", err)
			return
		}

		for _, stat := range priceRangeStats {
			t.Logf("价格区间: %s, 酒店数量: %d", stat.Range, stat.Count)
		}

		// 打印原始聚合结果
		aggJSON, _ := json.MarshalIndent(result.Aggregations, "", "  ")
		t.Logf("原始聚合结果: %s", string(aggJSON))
	})

	// 测试4：统计聚合查询
	t.Run("统计聚合查询", func(t *testing.T) {
		req := es_hotel.AggregationRequest{
			City: "上海",
		}

		searchReq, err := es_hotel.BuildStatsAggregationQuery(req)
		if err != nil {
			t.Errorf("构建统计聚合查询失败: %v", err)
			return
		}

		result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
		if err != nil {
			t.Errorf("执行统计聚合查询失败: %v", err)
			return
		}

		t.Logf("统计聚合查询结果: 总数=%d, 耗时=%dms", result.Total, result.Took)

		// 打印原始聚合结果
		aggJSON, _ := json.MarshalIndent(result.Aggregations, "", "  ")
		t.Logf("原始聚合结果: %s", string(aggJSON))
	})

	// 测试5：复杂嵌套聚合查询
	t.Run("复杂嵌套聚合查询", func(t *testing.T) {
		req := es_hotel.AggregationRequest{
			Keyword: "酒店",
		}

		searchReq, err := es_hotel.BuildNestedAggregationQuery(req)
		if err != nil {
			t.Errorf("构建嵌套聚合查询失败: %v", err)
			return
		}

		result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
		if err != nil {
			t.Errorf("执行嵌套聚合查询失败: %v", err)
			return
		}

		t.Logf("嵌套聚合查询结果: 总数=%d, 耗时=%dms", result.Total, result.Took)

		// 打印原始聚合结果
		aggJSON, _ := json.MarshalIndent(result.Aggregations, "", "  ")
		t.Logf("原始聚合结果: %s", string(aggJSON))
	})

	// 测试6：获取所有城市、品牌、星级的聚合查询
	t.Run("获取筛选选项聚合查询", func(t *testing.T) {
		// 构建筛选选项查询
		searchReq, err := es_hotel.BuildFilterOptionsQuery()
		if err != nil {
			t.Errorf("构建筛选选项查询失败: %v", err)
			return
		}

		// 执行查询
		result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
		if err != nil {
			t.Errorf("执行筛选选项查询失败: %v", err)
			return
		}

		t.Logf("筛选选项查询结果: 总数=%d, 耗时=%dms", result.Total, result.Took)

		// 解析筛选选项结果
		filterOptions, err := es_hotel.ParseFilterOptionsResult(result)
		if err != nil {
			t.Errorf("解析筛选选项结果失败: %v", err)
			return
		}

		// 输出城市选项
		t.Logf("城市选项数量: %d", len(filterOptions.Cities))
		for i, city := range filterOptions.Cities {
			if i < 10 { // 只显示前10个
				t.Logf("城市: %s, 酒店数量: %d", city.Label, city.Count)
			}
		}

		// 输出品牌选项
		t.Logf("品牌选项数量: %d", len(filterOptions.Brands))
		for i, brand := range filterOptions.Brands {
			if i < 10 { // 只显示前10个
				t.Logf("品牌: %s, 酒店数量: %d", brand.Label, brand.Count)
			}
		}

		// 输出星级选项
		t.Logf("星级选项数量: %d", len(filterOptions.StarNames))
		for _, star := range filterOptions.StarNames {
			t.Logf("星级: %s, 酒店数量: %d", star.Label, star.Count)
		}

		// 验证结果不为空
		if len(filterOptions.Cities) == 0 {
			t.Error("城市选项为空")
		}
		if len(filterOptions.Brands) == 0 {
			t.Error("品牌选项为空")
		}
		if len(filterOptions.StarNames) == 0 {
			t.Error("星级选项为空")
		}

		// 打印原始聚合结果（可选）
		aggJSON, _ := json.MarshalIndent(result.Aggregations, "", "  ")
		t.Logf("原始聚合结果: %s", string(aggJSON))

		// 测试便捷函数
		t.Run("测试GetFilterOptions便捷函数", func(t *testing.T) {
			filterOptions2, err := es_hotel.GetFilterOptions(typedESClient)
			if err != nil {
				t.Errorf("GetFilterOptions函数调用失败: %v", err)
				return
			}

			// 验证两种方式的结果一致
			if len(filterOptions2.Cities) != len(filterOptions.Cities) {
				t.Errorf("城市选项数量不一致: %d vs %d", len(filterOptions2.Cities), len(filterOptions.Cities))
			}
			if len(filterOptions2.Brands) != len(filterOptions.Brands) {
				t.Errorf("品牌选项数量不一致: %d vs %d", len(filterOptions2.Brands), len(filterOptions.Brands))
			}
			if len(filterOptions2.StarNames) != len(filterOptions.StarNames) {
				t.Errorf("星级选项数量不一致: %d vs %d", len(filterOptions2.StarNames), len(filterOptions.StarNames))
			}

			t.Logf("便捷函数测试通过，结果一致")
		})
	})
}

// TestSpecificAggregation 测试特定聚合查询
func TestSpecificAggregation(t *testing.T) {
	// 初始化ES客户端
	typedESClient, err := es_hotel.InitTypesElastic()
	if err != nil {
		t.Fatalf("初始化ES客户端失败: %v", err)
	}

	// 测试城市聚合
	req := es_hotel.AggregationRequest{}
	searchReq, err := es_hotel.BuildCityAggregationQuery(req)
	if err != nil {
		t.Fatalf("构建城市聚合查询失败: %v", err)
	}

	result, err := es_hotel.ExecuteAggregationQuery(searchReq, typedESClient)
	if err != nil {
		t.Fatalf("执行城市聚合查询失败: %v", err)
	}

	fmt.Printf("聚合查询结果:\n")
	fmt.Printf("总文档数: %d\n", result.Total)
	fmt.Printf("查询耗时: %dms\n", result.Took)
	fmt.Printf("聚合结果: %+v\n", result.Aggregations)
}
