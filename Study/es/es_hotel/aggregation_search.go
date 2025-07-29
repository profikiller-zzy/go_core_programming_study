package es_hotel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// AggregationRequest 聚合查询请求参数
type AggregationRequest struct {
	Keyword string `json:"keyword"` // 搜索关键词
	City    string `json:"city"`    // 城市过滤
}

// AggregationResult 聚合查询结果
type AggregationResult struct {
	Total       int64                  `json:"total"`       // 总记录数
	Took        int                    `json:"took"`        // 查询耗时(ms)
	Aggregations map[string]interface{} `json:"aggregations"` // 聚合结果
}

// CityStats 城市统计结果
type CityStats struct {
	City     string  `json:"city"`
	Count    int64   `json:"count"`
	AvgPrice float64 `json:"avg_price"`
	MaxPrice float64 `json:"max_price"`
	MinPrice float64 `json:"min_price"`
}

// BrandStats 品牌统计结果
type BrandStats struct {
	Brand    string  `json:"brand"`
	Count    int64   `json:"count"`
	AvgScore float64 `json:"avg_score"`
}

// PriceRangeStats 价格区间统计结果
type PriceRangeStats struct {
	Range string `json:"range"`
	Count int64  `json:"count"`
}

// 1. 基础聚合查询 - 按城市分组统计
func BuildCityAggregationQuery(req AggregationRequest) (*search.Request, error) {
	searchReq := &search.Request{
		Size: func() *int { size := 0; return &size }(), // 不返回文档，只要聚合结果
	}

	// 构建查询条件
	boolQuery := &types.BoolQuery{}
	if req.Keyword != "" {
		boolQuery.Must = []types.Query{
			{
				MultiMatch: &types.MultiMatchQuery{
					Query:  req.Keyword,
					Fields: []string{"name", "brand", "business"},
				},
			},
		}
	}

	searchReq.Query = &types.Query{
		Bool: boolQuery,
	}

	// 构建聚合
	aggregations := make(map[string]types.Aggregations)

	// 按城市分组聚合
	aggregations["cities"] = types.Aggregations{
		Terms: &types.TermsAggregation{
			Field: func() *string { field := "city"; return &field }(),
			Size:  func() *int { size := 10; return &size }(),
		},
		// 嵌套聚合：计算每个城市的平均价格
		Aggregations: map[string]types.Aggregations{
			"avg_price": {
				Avg: &types.AverageAggregation{
					Field: func() *string { field := "price"; return &field }(),
				},
			},
			"max_price": {
				Max: &types.MaxAggregation{
					Field: func() *string { field := "price"; return &field }(),
				},
			},
			"min_price": {
				Min: &types.MinAggregation{
					Field: func() *string { field := "price"; return &field }(),
				},
			},
		},
	}

	searchReq.Aggregations = aggregations
	return searchReq, nil
}

// 2. 多维度聚合查询 - 品牌和星级统计
func BuildMultiDimensionAggregationQuery(req AggregationRequest) (*search.Request, error) {
	searchReq := &search.Request{
		Size: func() *int { size := 0; return &size }(),
	}

	// 构建查询条件
	boolQuery := &types.BoolQuery{}
	if req.City != "" {
		boolQuery.Filter = []types.Query{
			{
				Term: map[string]types.TermQuery{
					"city": {
						Value: req.City,
					},
				},
			},
		}
	}

	searchReq.Query = &types.Query{
		Bool: boolQuery,
	}

	// 构建多维度聚合
	aggregations := make(map[string]types.Aggregations)

	// 品牌聚合
	aggregations["brands"] = types.Aggregations{
		Terms: &types.TermsAggregation{
			Field: func() *string { field := "brand"; return &field }(),
			Size:  func() *int { size := 20; return &size }(),
		},
		// 嵌套聚合：每个品牌的平均评分
		Aggregations: map[string]types.Aggregations{
			"avg_score": {
				Avg: &types.AverageAggregation{
					Field: func() *string { field := "score"; return &field }(),
				},
			},
			// 再嵌套：每个品牌下的星级分布
			"star_distribution": {
				Terms: &types.TermsAggregation{
					Field: func() *string { field := "starName"; return &field }(),
					Size:  func() *int { size := 10; return &size }(),
				},
			},
		},
	}

	// 星级聚合
	aggregations["star_levels"] = types.Aggregations{
		Terms: &types.TermsAggregation{
			Field: func() *string { field := "starName"; return &field }(),
			Size:  func() *int { size := 10; return &size }(),
		},
	}

	searchReq.Aggregations = aggregations
	return searchReq, nil
}

// 3. 范围聚合查询 - 价格区间统计
func BuildRangeAggregationQuery(req AggregationRequest) (*search.Request, error) {
	searchReq := &search.Request{
		Size: func() *int { size := 0; return &size }(),
	}

	// 构建查询条件
	boolQuery := &types.BoolQuery{}
	if req.Keyword != "" {
		boolQuery.Must = []types.Query{
			{
				Match: map[string]types.MatchQuery{
					"name": {
						Query: req.Keyword,
					},
				},
			},
		}
	}

	searchReq.Query = &types.Query{
		Bool: boolQuery,
	}

	// 构建范围聚合
	aggregations := make(map[string]types.Aggregations)

	// 价格区间聚合
	aggregations["price_ranges"] = types.Aggregations{
		Range: &types.RangeAggregation{
			Field: func() *string { field := "price"; return &field }(),
			Ranges: []types.AggregationRange{
				{
					Key: func() *string { key := "0-200"; return &key }(),
					To:  func() *types.Float64 { to := types.Float64(200); return &to }(),
				},
				{
					Key:  func() *string { key := "200-500"; return &key }(),
					From: func() *types.Float64 { from := types.Float64(200); return &from }(),
					To:   func() *types.Float64 { to := types.Float64(500); return &to }(),
				},
				{
					Key:  func() *string { key := "500-1000"; return &key }(),
					From: func() *types.Float64 { from := types.Float64(500); return &from }(),
					To:   func() *types.Float64 { to := types.Float64(1000); return &to }(),
				},
				{
					Key:  func() *string { key := "1000+"; return &key }(),
					From: func() *types.Float64 { from := types.Float64(1000); return &from }(),
				},
			},
		},
	}

	// 评分区间聚合
	aggregations["score_ranges"] = types.Aggregations{
		Range: &types.RangeAggregation{
			Field: func() *string { field := "score"; return &field }(),
			Ranges: []types.AggregationRange{
				{
					Key: func() *string { key := "1-2分"; return &key }(),
					To:  func() *types.Float64 { to := types.Float64(2); return &to }(),
				},
				{
					Key:  func() *string { key := "2-3分"; return &key }(),
					From: func() *types.Float64 { from := types.Float64(2); return &from }(),
					To:   func() *types.Float64 { to := types.Float64(3); return &to }(),
				},
				{
					Key:  func() *string { key := "3-4分"; return &key }(),
					From: func() *types.Float64 { from := types.Float64(3); return &from }(),
					To:   func() *types.Float64 { to := types.Float64(4); return &to }(),
				},
				{
					Key:  func() *string { key := "4-5分"; return &key }(),
					From: func() *types.Float64 { from := types.Float64(4); return &from }(),
					To:   func() *types.Float64 { to := types.Float64(5); return &to }(),
				},
			},
		},
	}

	searchReq.Aggregations = aggregations
	return searchReq, nil
}

// 4. 统计聚合查询 - 价格和评分统计
func BuildStatsAggregationQuery(req AggregationRequest) (*search.Request, error) {
	searchReq := &search.Request{
		Size: func() *int { size := 0; return &size }(),
	}

	// 构建查询条件
	boolQuery := &types.BoolQuery{}
	if req.City != "" {
		boolQuery.Filter = []types.Query{
			{
				Term: map[string]types.TermQuery{
					"city": {
						Value: req.City,
					},
				},
			},
		}
	}

	searchReq.Query = &types.Query{
		Bool: boolQuery,
	}

	// 构建统计聚合
	aggregations := make(map[string]types.Aggregations)

	// 价格统计
	aggregations["price_stats"] = types.Aggregations{
		Stats: &types.StatsAggregation{
			Field: func() *string { field := "price"; return &field }(),
		},
	}

	// 评分统计
	aggregations["score_stats"] = types.Aggregations{
		Stats: &types.StatsAggregation{
			Field: func() *string { field := "score"; return &field }(),
		},
	}

	// 扩展统计（包含方差、标准差等）
	aggregations["price_extended_stats"] = types.Aggregations{
		ExtendedStats: &types.ExtendedStatsAggregation{
			Field: func() *string { field := "price"; return &field }(),
		},
	}

	searchReq.Aggregations = aggregations
	return searchReq, nil
}

// 5. 复杂嵌套聚合查询 - 城市->品牌->星级的三层嵌套
func BuildNestedAggregationQuery(req AggregationRequest) (*search.Request, error) {
	searchReq := &search.Request{
		Size: func() *int { size := 0; return &size }(),
	}

	// 构建查询条件
	boolQuery := &types.BoolQuery{}
	if req.Keyword != "" {
		boolQuery.Must = []types.Query{
			{
				MultiMatch: &types.MultiMatchQuery{
					Query:  req.Keyword,
					Fields: []string{"name", "brand", "business"},
				},
			},
		}
	}

	searchReq.Query = &types.Query{
		Bool: boolQuery,
	}

	// 构建三层嵌套聚合：城市 -> 品牌 -> 星级
	aggregations := make(map[string]types.Aggregations)

	aggregations["cities"] = types.Aggregations{
		Terms: &types.TermsAggregation{
			Field: func() *string { field := "city"; return &field }(),
			Size:  func() *int { size := 10; return &size }(),
		},
		// 第二层：品牌聚合
		Aggregations: map[string]types.Aggregations{
			"brands": {
				Terms: &types.TermsAggregation{
					Field: func() *string { field := "brand"; return &field }(),
					Size:  func() *int { size := 10; return &size }(),
				},
				// 第三层：星级聚合
				Aggregations: map[string]types.Aggregations{
					"star_levels": {
						Terms: &types.TermsAggregation{
							Field: func() *string { field := "starName"; return &field }(),
							Size:  func() *int { size := 5; return &size }(),
						},
						// 第四层：统计信息
						Aggregations: map[string]types.Aggregations{
							"avg_price": {
								Avg: &types.AverageAggregation{
									Field: func() *string { field := "price"; return &field }(),
								},
							},
							"avg_score": {
								Avg: &types.AverageAggregation{
									Field: func() *string { field := "score"; return &field }(),
								},
							},
						},
					},
				},
			},
			// 城市级别的统计
			"city_avg_price": {
				Avg: &types.AverageAggregation{
					Field: func() *string { field := "price"; return &field }(),
				},
			},
			"city_avg_score": {
				Avg: &types.AverageAggregation{
					Field: func() *string { field := "score"; return &field }(),
				},
			},
		},
	}

	searchReq.Aggregations = aggregations
	return searchReq, nil
}

// 执行聚合查询的通用函数
func ExecuteAggregationQuery(searchReq *search.Request, typedESClient *elasticsearch.TypedClient) (*AggregationResult, error) {
	// 执行搜索
	response, err := typedESClient.Search().
		Index(HotelIndex).
		Request(searchReq).
		Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("执行聚合查询失败: %v", err)
	}

	// 解析聚合结果
	aggregations := make(map[string]interface{})
	if response.Aggregations != nil {
		// 将聚合结果转换为map
		aggBytes, err := json.Marshal(response.Aggregations)
		if err != nil {
			return nil, fmt.Errorf("序列化聚合结果失败: %v", err)
		}

		err = json.Unmarshal(aggBytes, &aggregations)
		if err != nil {
			return nil, fmt.Errorf("反序列化聚合结果失败: %v", err)
		}
	}

	return &AggregationResult{
		Total:        response.Hits.Total.Value,
		Took:         int(response.Took),
		Aggregations: aggregations,
	}, nil
}

// 解析城市聚合结果的辅助函数
func ParseCityAggregationResult(result *AggregationResult) ([]CityStats, error) {
	citiesAgg, ok := result.Aggregations["cities"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("城市聚合结果格式错误")
	}

	buckets, ok := citiesAgg["buckets"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("城市聚合桶格式错误")
	}

	var cityStats []CityStats
	for _, bucket := range buckets {
		bucketMap, ok := bucket.(map[string]interface{})
		if !ok {
			continue
		}

		city, _ := bucketMap["key"].(string)
		count, _ := bucketMap["doc_count"].(float64)

		// 解析嵌套聚合结果
		avgPriceAgg, _ := bucketMap["avg_price"].(map[string]interface{})
		maxPriceAgg, _ := bucketMap["max_price"].(map[string]interface{})
		minPriceAgg, _ := bucketMap["min_price"].(map[string]interface{})

		avgPrice, _ := avgPriceAgg["value"].(float64)
		maxPrice, _ := maxPriceAgg["value"].(float64)
		minPrice, _ := minPriceAgg["value"].(float64)

		cityStats = append(cityStats, CityStats{
			City:     city,
			Count:    int64(count),
			AvgPrice: avgPrice,
			MaxPrice: maxPrice,
			MinPrice: minPrice,
		})
	}

	return cityStats, nil
}

// 解析品牌聚合结果的辅助函数
func ParseBrandAggregationResult(result *AggregationResult) ([]BrandStats, error) {
	brandsAgg, ok := result.Aggregations["brands"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("品牌聚合结果格式错误")
	}

	buckets, ok := brandsAgg["buckets"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("品牌聚合桶格式错误")
	}

	var brandStats []BrandStats
	for _, bucket := range buckets {
		bucketMap, ok := bucket.(map[string]interface{})
		if !ok {
			continue
		}

		brand, _ := bucketMap["key"].(string)
		count, _ := bucketMap["doc_count"].(float64)

		// 解析平均评分
		avgScoreAgg, _ := bucketMap["avg_score"].(map[string]interface{})
		avgScore, _ := avgScoreAgg["value"].(float64)

		brandStats = append(brandStats, BrandStats{
			Brand:    brand,
			Count:    int64(count),
			AvgScore: avgScore,
		})
	}

	return brandStats, nil
}

// 解析价格区间聚合结果的辅助函数
func ParsePriceRangeAggregationResult(result *AggregationResult) ([]PriceRangeStats, error) {
	priceRangesAgg, ok := result.Aggregations["price_ranges"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("价格区间聚合结果格式错误")
	}

	buckets, ok := priceRangesAgg["buckets"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("价格区间聚合桶格式错误")
	}

	var priceRangeStats []PriceRangeStats
	for _, bucket := range buckets {
		bucketMap, ok := bucket.(map[string]interface{})
		if !ok {
			continue
		}

		rangeKey, _ := bucketMap["key"].(string)
		count, _ := bucketMap["doc_count"].(float64)

		priceRangeStats = append(priceRangeStats, PriceRangeStats{
			Range: rangeKey,
			Count: int64(count),
		})
	}

	return priceRangeStats, nil
}