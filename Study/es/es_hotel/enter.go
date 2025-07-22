package es_hotel

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ElasticClient *elasticsearch.Client
	HotelIndex    string = "hotel"
)

// InitElastic 初始化ES客户端
func InitElastic() error {
	var err error
	ElasticClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://8.155.39.153:9200"}, // 根据你的ES地址修改
		// Username:  "elastic",                        // 如果需要认证，取消注释
		// Password:  "your_password",                  // 如果需要认证，取消注释
	})
	if err != nil {
		return fmt.Errorf("创建ES客户端失败: %v", err)
	}

	// 测试连接
	res, err := ElasticClient.Info()
	if err != nil {
		return fmt.Errorf("ES连接测试失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("ES连接错误: %s", res.String())
	}

	fmt.Println("ES连接成功")
	return nil
}

// InitTypesElastic 初始化ES客户端
func InitTypesElastic() (*elasticsearch.TypedClient, error) {
	var err error
	typedESClient, err := elasticsearch.NewTypedClient(
		elasticsearch.Config{
			Addresses: []string{"http://8.155.39.153:9200"}, // 根据你的ES地址修改
			// Username:  "elastic",                        // 如果需要认证，取消注释
			// Password:  "your_password",                  // 如果需要认证，取消注释
		})
	if err != nil {
		return nil, fmt.Errorf("创建ES客户端失败: %v", err)
	}
	return typedESClient, nil
}

// InitMySQL 初始化MySQL连接
func InitMySQL() (*gorm.DB, error) {
	// 根据你的MySQL配置修改连接字符串
	dsn := "root:01312934a@tcp(localhost:3306)/mysql_study?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("连接MySQL失败: %v", err)
	}

	fmt.Println("MySQL连接成功")
	return mysqlDB, nil
}
