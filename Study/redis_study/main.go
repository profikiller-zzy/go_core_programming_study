package main

import (
	"github.com/redis/go-redis/v9"
	"go_core_programming/Study/redis_study/hyperloglog"
)

func main() {
	redisClient := InitRedis()
	hyperloglog.TestHyperloglog(redisClient)
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "8.155.34.188:6379", // Redis server address
		Password: "38163816aA",
		DB:       0, // Default DB
	})
}
