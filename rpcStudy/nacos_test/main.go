package main

import (
	"fmt"
	"log"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func main() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "c9f609db-3402-4d70-bd0c-fef7f5f599f8", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "8.155.39.153",
			Port:   8848,
			Scheme: "http",
		},
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to create config client: %v", err))
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.yaml",
		Group:  "dev",
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to get config: %v", err))
	}
	fmt.Println("Config content:", content)

	var config Config
	err = yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// 输出解析结果
	fmt.Printf("Name: %s\n", config.Name)
	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %s\n", config.Port)
}
