package es_hotel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// GetHotelsFromMySQL 从MySQL获取所有酒店数据
func GetHotelsFromMySQL(ctx context.Context, db *gorm.DB) ([]Hotel, error) {
	var hotels []Hotel
	// 执行查询
	err := db.WithContext(ctx).Model(&Hotel{}).Find(&hotels).Error
	if err != nil {
		return nil, fmt.Errorf("查询酒店数据失败: %v", err)
	}

	fmt.Printf("从MySQL获取到 %d 条酒店数据\n", len(hotels))
	return hotels, nil
}

// CreateHotelIndex 创建酒店索引（如果不存在）
func CreateHotelIndex() error {
	// 检查索引是否存在
	res, err := ElasticClient.Indices.Exists([]string{HotelIndex})
	if err != nil {
		return fmt.Errorf("检查索引存在性失败: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Printf("索引 %s 已存在\n", HotelIndex)
		return nil
	}

	// 创建索引映射
	mapping :=
		`{
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"name": {
				"type": "text",
				"analyzer": "ik_smart",
				"search_analyzer": "ik_smart",
				"copy_to": "all"
			},
			"address": {
				"type": "keyword",
				"index": false
			},
			"price": {
				"type": "integer"
			},
			"score": {
				"type": "integer"
			},
			"brand": {
				"type": "keyword",
				"copy_to": "all"
			},
			"city": {
				"type": "keyword"
			},
			"starName": {
				"type": "keyword"
			},
			"business": {
				"type": "keyword",
				"copy_to": "all"
			},
			"location": {
				"type": "geo_point"
			},
			"pic": {
				"type": "keyword",
				"index": false
			},
			"all": {
				"type": "text",
				"analyzer": "ik_smart",
				"search_analyzer": "ik_smart"
			},
			"suggestion": {
				"type": "completion",
				"analyzer": "standard"
			}
		}
	}
}`

	req := esapi.IndicesCreateRequest{
		Index: HotelIndex,
		Body:  strings.NewReader(mapping),
	}

	res, err = req.Do(context.Background(), ElasticClient)
	if err != nil {
		return fmt.Errorf("创建索引失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("创建索引响应错误: %s", res.String())
	}

	fmt.Printf("索引 %s 创建成功\n", HotelIndex)
	return nil
}

// BulkIndexHotels 批量导入酒店数据到ES
func BulkIndexHotels(hotels []Hotel) error {
	const batchSize = 100 // 每批处理100条数据

	for i := 0; i < len(hotels); i += batchSize {
		end := i + batchSize
		if end > len(hotels) {
			end = len(hotels)
		}

		batch := hotels[i:end]
		if err := indexBatch(batch); err != nil {
			return fmt.Errorf("批量导入第 %d-%d 条数据失败: %v", i+1, end, err)
		}

		fmt.Printf("成功导入第 %d-%d 条数据\n", i+1, end)
	}

	return nil
}

// indexBatch 导入一批数据
func indexBatch(hotels []Hotel) error {
	var buf bytes.Buffer

	for _, hotel := range hotels {
		// 添加操作元数据
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": HotelIndex,
				"_id":    fmt.Sprintf("%d", hotel.ID),
			},
		}
		metaJSON, _ := json.Marshal(meta)
		buf.Write(metaJSON)
		buf.WriteByte('\n')

		// 添加地理位置字段
		hotelDoc := struct {
			Hotel
			Location map[string]interface{} `json:"location"`
		}{
			Hotel: hotel,
			Location: map[string]interface{}{
				"lat": hotel.Latitude,
				"lon": hotel.Longitude,
			},
		}

		// 添加文档数据
		docJSON, err := json.Marshal(hotelDoc)
		if err != nil {
			return fmt.Errorf("序列化酒店数据失败: %v", err)
		}
		buf.Write(docJSON)
		buf.WriteByte('\n')
	}

	// 执行批量请求
	req := esapi.BulkRequest{
		Body:    strings.NewReader(buf.String()),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), ElasticClient)
	if err != nil {
		return fmt.Errorf("批量请求失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("批量请求响应错误: %s", res.String())
	}

	// 解析响应检查是否有错误
	var bulkResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&bulkResponse); err != nil {
		return fmt.Errorf("解析批量响应失败: %v", err)
	}

	if errors, ok := bulkResponse["errors"].(bool); ok && errors {
		return fmt.Errorf("批量操作存在错误")
	}

	return nil
}
