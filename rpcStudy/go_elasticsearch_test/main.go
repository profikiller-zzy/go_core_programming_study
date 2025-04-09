package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
)

type User struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
	Desc string `json:"desc"`
}

// 解析时使用结构体
var response struct {
	Index   string `json:"_index"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Source  User   `json:"_source"`
}

func main() {
	// 初始化一个连接es的客户端
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://8.155.39.153:9200",
		},
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")

	// 索引一条数据
	user := User{
		Name: "zjy",
		Age:  18,
		Desc: "always be happy, and have a beautiful girl friend",
	}
	data, err := json.Marshal(user)
	if err != nil {
		logger.Error(err)
		return
	}
	_, err = esClient.Index("user", bytes.NewReader(data))
	if err != nil {
		return
	}

	// 按照id get一条数据
	res, err := esClient.Get("user", "IIcSopUBYQtX7iavlwVm")
	if err != nil {
		fmt.Println(err)
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("解析失败: %s", err)
	}
	fmt.Println(response.Source)

	// 使用match查询
	query := `{
		"query": {
			"match_all": {}
		},
		"size": 10,    // 返回文档数
		"from": 0      // 分页起始位置
	}`
	// 执行 Search 请求
	res, err = esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("user"), // 替换为索引名
		esClient.Search.WithBody(strings.NewReader(query)),
		esClient.Search.WithTrackTotalHits(true), // 确保返回总命中数
		esClient.Search.WithPretty(),             // 美化响应
	)
	if err != nil {
		log.Fatalf("Search error: %s", err)
	}
	defer res.Body.Close()
	// 检查 HTTP 状态码
	if res.IsError() {
		log.Fatalf("Search failed: %s", res.String())
	}
	// 解析响应体
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing response: %s", err)
	}
	// 提取关键数据
	total := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	fmt.Printf("Total documents: %d\n", int(total))
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		id := hit.(map[string]interface{})["_id"].(string)
		score := hit.(map[string]interface{})["_score"].(float64)

		jsonData, err := json.Marshal(source)
		user := User{}
		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("ID: %s | Score: %.2f\n%v\n", id, score, user)
	}

	indexName := "goods"
	mappings :=
		`{
		"mappings": {
			"properties": {
				"name": {
					"type": "text",
					"analyzer": "ik_max_word"
				},
				"id": {
					"type": "integer"
				}
			}
		}
	}`
	// 创建请求
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(mappings),
	}
	// 执行请求
	res, err = req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// 解析响应
	if res.IsError() {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			log.Fatalf("Error parsing error response: %s", err)
		}
		log.Fatalf("Failed to create index: [%s] %s",
			res.Status(),
			errorResponse["error"].(map[string]interface{})["reason"],
		)
	}
	// 成功响应
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing response: %s", err)
	}
	fmt.Printf("Index '%s' created: %v\n", indexName, response["acknowledged"])
}
