package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

var currentIP = "183.241.254.165"

func registerService(name string, id string, tags []string, address string, port int) {
	cfg := api.DefaultConfig()
	cfg.Address = "8.155.39.153:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 创建注册对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://8.155.39.153:8021/health",
		Interval:                       "5s",  // 健康检查间隔
		Timeout:                        "5s",  // 超时时间，超过这个时间没有返回则认为是不健康
		DeregisterCriticalServiceAfter: "10s", // 多久后注销服务
	}
	reg := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
		Check:   check,
	}

	// 注册服务
	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}
}

func getService() {
	cfg := api.DefaultConfig()
	cfg.Address = "8.155.39.153:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, value := range data {
		fmt.Println(key, value)
	}
}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "8.155.39.153:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(`Service == "rpcService"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	//registerService("rpcService", "rpcService", []string{"rpc", "mxshop"}, currentIP, 50051)
	getService()
	FilterService()
}
