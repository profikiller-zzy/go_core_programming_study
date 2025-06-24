package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	redisClient := InitRedis()
	ctx := context.Background()
	if _, err := redisClient.Set(ctx, "key", "value", 300*time.Second).Result(); err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["name"] = "黑马产品"
	data["price"] = 100
	data["stock"] = 1000
	vals, err := redisClient.HMSet(ctx, "heima:product:1", data).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

	value, err := redisClient.HGetAll(ctx, "heima:product:1").Result()
	if err != nil {
		panic(err)
	}
	for key, v := range value {
		fmt.Println(key, v)
	}
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "8.155.34.188:6379", // Redis server address
		Password: "38163816aA",
		DB:       0, // Default DB
	})
}
