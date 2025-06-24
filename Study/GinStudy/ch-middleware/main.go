package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Logger 自定义中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Abort()
		t := time.Now()

		// 在请求前设置变量
		c.Set("example", "12345")

		c.Next() // 处理后续中间件和路由

		// 请求后计算耗时
		latency := time.Since(t)
		log.Print("请求耗时:", latency)

		// 获取响应状态码
		status := c.Writer.Status()
		log.Println("状态码:", status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger()) // 应用自定义中间件

	r.GET("/test", func(c *gin.Context) {
		// 获取中间件设置的变量
		example := c.MustGet("example").(string)
		log.Println("中间件传递的值:", example) // 输出 "12345"
	})

	r.Run(":8080")
}
