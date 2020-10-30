package main

import (
	"mycode/nightingale/base"
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

	// r.GET("/test", service.GetMsg)

	r.GET("/tt", BaseFun(service.GetMsg))
	r.GET("/mm", BaseFun(service.RuleFindAll))
	r.Run()
}

func BaseFun(op func(c *gin.Context) base.Any) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := op(c)
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
