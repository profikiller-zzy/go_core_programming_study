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

	success, err := redisClient.SetNX(ctx, "lock", "locked", 10*time.Second).Result()
	if err != nil {
		panic(err)
	}
	if success {
		fmt.Println("locked")
	} else {
		fmt.Println("lock already exists")
	}

	code, err := redisClient.Del(ctx, "lock").Result()
	if err != nil {
		panic(err)
	}
	if code > 0 {
		fmt.Println("lock released")
	} else {
		fmt.Println("lock not found")
	}
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "8.155.34.188:6379", // Redis server address
		Password: "38163816aA",
		DB:       0, // Default DB
	})
}
