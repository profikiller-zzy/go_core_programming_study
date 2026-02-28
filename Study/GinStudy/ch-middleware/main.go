package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// TimingMiddleware 统计接口执行时间的中间件
func TimingMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 记录请求开始时间
		start := time.Now()

		// 获取请求信息
		method := string(c.Method())
		path := string(c.Path())
		clientIP := c.ClientIP()

		log.Printf("[%s] %s %s - 开始处理请求", method, path, clientIP)

		// 继续处理后续中间件和路由处理函数
		c.Next(ctx)

		// 计算执行时间
		duration := time.Since(start)
		status := c.Response.StatusCode()

		// 记录请求完成信息
		log.Printf("[%s] %s %s - 状态码: %d, 执行时间: %v",
			method, path, clientIP, status, duration)

		// 可以根据执行时间进行不同的处理
		if duration > 1*time.Second {
			log.Printf("⚠️  慢请求警告: %s %s 执行时间 %v 超过 1 秒", method, path, duration)
		}
	}
}

func main() {
	// 创建 Hertz 服务器
	h := server.Default(server.WithHostPorts(":8080"))

	// 使用执行时间统计中间件
	h.Use(TimingMiddleware())
	// 或者使用高级版本

	// 定义路由
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		// 模拟一些处理时间
		time.Sleep(100 * time.Millisecond)
		c.JSON(200, map[string]interface{}{
			"message":   "Hello Hertz!",
			"timestamp": time.Now(),
		})
	})

	h.Spin()
}
