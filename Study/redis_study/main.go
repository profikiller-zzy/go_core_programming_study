package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/inner/uuid"
	"github.com/redis/go-redis/v9"

	"go_core_programming/Study/redis_study/redis_lock"
)

func main() {
	redisClient := InitRedis()
	uid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error generating UUID:", err)
		return
	}
	redisLock := redis_lock.NewSimpleRedisLock("order", uid.String(), redisClient)
	ok, err := redisLock.TryLock("123456", 10)
	if err != nil {
		fmt.Println("Error trying to lock:", err)
		return
	}
	if !ok {
		fmt.Println("Failed to acquire lock")
		return
	}
	time.Sleep(time.Second * 2)

	err = redisLock.Unlock("123456")
	fmt.Println("成功释放锁")
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "8.155.34.188:6379", // Redis server address
		Password: "38163816aA",
		DB:       0, // Default DB
	})
}
