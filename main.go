package main

import (
	"mycode/nightingale/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messgae": "pong",
		})
	})

	r.GET("/test", service.GetMsg)

	r.Run()
}
