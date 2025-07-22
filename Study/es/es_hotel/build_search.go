package es_hotel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// SearchRequest 搜索请求参数
type SearchRequest struct {
	// 基础搜索参数
	Keyword string `json:"keyword"` // 搜索关键词
	Page    int    `json:"page"`    // 页码，从1开始
	Size    int    `json:"size"`    // 每页大小

	// 过滤条件
	City     string `json:"city"`     // 城市过滤
	Brand    string `json:"brand"`    // 品牌过滤
	StarName string `json:"starName"` // 星级过滤
	Business string `json:"business"` // 商圈过滤

	// 价格范围过滤
	MinPrice float64 `json:"minPrice"` // 最低价格
	MaxPrice float64 `json:"maxPrice"` // 最高价格

	// 评分范围过滤
	MinScore float64 `json:"minScore"` // 最低评分
	MaxScore float64 `json:"maxScore"` // 最高评分

	// 地理位置过滤
	Latitude  float64 `json:"latitude"`  // 纬度
	Longitude float64 `json:"longitude"` // 经度
	Distance  string  `json:"distance"`  // 距离范围，如"5km"

	// 排序参数
	SortBy    string `json:"sortBy"`    // 排序字段：price, score, _geo_distance
	SortOrder string `json:"sortOrder"` // 排序方向：asc, desc

	// 高亮参数
	Highlight bool `json:"highlight"` // 是否启用高亮
}

// BuildSearchRequest 构建ES搜索请求（使用官方SDK结构体）
func BuildSearchRequest(req SearchRequest) (*search.Request, error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	if req.Size > 100 {
		req.Size = 100 // 限制最大页面大小
	}

	from := (req.Page - 1) * req.Size
	size := req.Size

	// 创建搜索请求
	searchReq := &search.Request{
		From: &from,
		Size: &size,
	}

	// 构建bool查询
	boolQuery := &types.BoolQuery{}

	// 构建must查询（搜索条件）
	var mustQueries []types.Query
	if strings.TrimSpace(req.Keyword) != "" {
		// 使用multi_match查询多个字段
		multiMatch := &types.MultiMatchQuery{
			Query:  req.Keyword,
			Fields: []string{"all", "name^2", "brand", "business"},
		}
		mustQueries = append(mustQueries, types.Query{
			MultiMatch: multiMatch,
		})
	}
	if len(mustQueries) > 0 {
		boolQuery.Must = mustQueries
	} else {
		// 如果没有搜索条件，使用match_all
		boolQuery.Must = []types.Query{
			{
				MatchAll: &types.MatchAllQuery{},
			},
		}
	}

	// 构建filter查询（过滤条件）
	var filterQueries []types.Query

	// 城市过滤
	if strings.TrimSpace(req.City) != "" {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"city": {
					Value: req.City,
				},
			},
		})
	}

	// 品牌过滤
	if strings.TrimSpace(req.Brand) != "" {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"brand": {
					Value: req.Brand,
				},
			},
		})
	}

	// 星级过滤
	if strings.TrimSpace(req.StarName) != "" {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"starName": {
					Value: req.StarName,
				},
			},
		})
	}

	// 商圈过滤
	if strings.TrimSpace(req.Business) != "" {
		filterQueries = append(filterQueries, types.Query{
			Term: map[string]types.TermQuery{
				"business": {
					Value: req.Business,
				},
			},
		})
	}

	// 价格范围过滤
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		rangeQuery := types.NumberRangeQuery{}
		if req.MinPrice > 0 {
			gte := types.Float64(req.MinPrice)
			rangeQuery.Gte = &gte
		}
		if req.MaxPrice > 0 {
			lte := types.Float64(req.MaxPrice)
			rangeQuery.Lte = &lte
		}
		filterQueries = append(filterQueries, types.Query{
			Range: map[string]types.RangeQuery{
				"price": rangeQuery,
			},
		})
	}

	// 评分范围过滤
	if req.MinScore > 0 || req.MaxScore > 0 {
		rangeQuery := types.NumberRangeQuery{}
		if req.MinScore > 0 {
			gte := types.Float64(req.MinScore)
			rangeQuery.Gte = &gte
		}
		if req.MaxScore > 0 {
			lte := types.Float64(req.MaxScore)
			rangeQuery.Lte = &lte
		}
		filterQueries = append(filterQueries,
			types.Query{
				Range: map[string]types.RangeQuery{
					"score": rangeQuery,
				},
			})
	}

	// 地理位置过滤
	if req.Latitude != 0 && req.Longitude != 0 && strings.TrimSpace(req.Distance) != "" {
		geoDistanceQuery := &types.GeoDistanceQuery{
			Distance: req.Distance,
			GeoDistanceQuery: map[string]types.GeoLocation{
				"location": types.LatLonGeoLocation{
					Lat: types.Float64(req.Latitude),
					Lon: types.Float64(req.Longitude),
				},
			},
		}
		filterQueries = append(filterQueries,
			types.Query{
				GeoDistance: geoDistanceQuery,
			})
	}

	// 添加filter到bool查询
	if len(filterQueries) > 0 {
		boolQuery.Filter = filterQueries
	}

	// 设置查询
	searchReq.Query = &types.Query{
		Bool: boolQuery,
	}

	// 构建排序
	if req.SortBy != "" {
		switch req.SortBy {
		case "price":
			searchReq.Sort = []types.SortCombinations{
				types.SortOptions{
					SortOptions: map[string]types.FieldSort{
						"price": {
							Order: &sortorder.Asc,
						},
					},
				},
			}
		case "score":
			searchReq.Sort = []types.SortCombinations{
				types.SortOptions{
					SortOptions: map[string]types.FieldSort{
						"score": {
							Order: &sortorder.Asc,
						},
					},
				},
			}
		case "_geo_distance":
			searchReq.Sort = []types.SortCombinations{
				types.SortOptions{
					GeoDistance_: &types.GeoDistanceSort{
						GeoDistanceSort: map[string][]types.GeoLocation{
							"location": {
								types.LatLonGeoLocation{
									Lat: types.Float64(req.Latitude),
									Lon: types.Float64(req.Longitude),
								},
							},
						},
						Order: &sortorder.SortOrder{
							Name: "asc", // 使用升序排序
						},
						Unit: &distanceunit.DistanceUnit{
							Name: "km", // 使用千米作为距离单位
						},
					},
				},
			}
		default:
			searchReq.Sort = nil
		}
	}

	// 构建高亮
	if req.Highlight {
		fragmentSize := 100
		numberOfFragments := 1
		allFragmentSize := 150
		allNumberOfFragments := 2

		highlight := &types.Highlight{
			PreTags:  []string{"<em>"},
			PostTags: []string{"</em>"},
			Fields: map[string]types.HighlightField{
				"name": {
					FragmentSize:      &fragmentSize,
					NumberOfFragments: &numberOfFragments,
				},
				"brand": {
					FragmentSize:      &fragmentSize,
					NumberOfFragments: &numberOfFragments,
				},
				"business": {
					FragmentSize:      &fragmentSize,
					NumberOfFragments: &numberOfFragments,
				},
				"all": {
					FragmentSize:      &allFragmentSize,
					NumberOfFragments: &allNumberOfFragments,
				},
			},
		}
		searchReq.Highlight = highlight
	}

	return searchReq, nil
}

// SearchResult 格式化的搜索结果
type SearchResult struct {
	Total  int64   `json:"total"`  // 总记录数
	Page   int     `json:"page"`   // 当前页码
	Size   int     `json:"size"`   // 每页大小
	Pages  int     `json:"pages"`  // 总页数
	Hotels []Hotel `json:"hotels"` // 酒店列表
	Took   int     `json:"took"`   // 查询耗时(ms)
}

// ParseSearchResponse 解析ES搜索响应（使用官方SDK响应结构体）
func ParseSearchResponse(response *search.Response, page, size int) (*SearchResult, error) {
	if response == nil {
		return nil, fmt.Errorf("搜索响应为空")
	}

	// 提取酒店数据
	hotels := make([]Hotel, 0, len(response.Hits.Hits))
	for _, hit := range response.Hits.Hits {
		// 解析_source到Hotel结构体
		var hotel Hotel
		if hit.Source_ != nil {
			// 将hit.Source_（json.RawMessage）解析为Hotel结构体
			err := json.Unmarshal(hit.Source_, &hotel)
			if err != nil {
				return nil, fmt.Errorf("解析酒店数据失败: %v", err)
			}
		}
		hotels = append(hotels, hotel)
	}

	// 计算总页数
	total := response.Hits.Total.Value
	pages := int(total) / size
	if int(total)%size > 0 {
		pages++
	}

	return &SearchResult{
		Total:  total,
		Page:   page,
		Size:   size,
		Pages:  pages,
		Hotels: hotels,
		Took:   int(response.Took),
	}, nil
}

// ExecuteSearch 执行搜索（使用官方SDK）
func ExecuteSearch(req SearchRequest, typedESClient *elasticsearch.TypedClient) (*SearchResult, error) {
	// 构建搜索请求
	searchReq, err := BuildSearchRequest(req)
	if err != nil {
		return nil, fmt.Errorf("构建搜索请求失败: %v", err)
	}

	// 执行搜索
	response, err := typedESClient.Search().
		Index(HotelIndex).
		Request(searchReq).
		Do(context.Background())

	if err != nil {
		return nil, fmt.Errorf("执行搜索失败: %v", err)
	}

	// 解析响应
	return ParseSearchResponse(response, req.Page, req.Size)
}
