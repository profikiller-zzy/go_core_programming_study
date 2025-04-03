package main

import "github.com/gin-gonic/gin"

func npsReturn(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "连接成功",
	})
}

func main() {
	r := gin.Default()
	r.GET("/nps", npsReturn)
	r.Run(":8091")
}
