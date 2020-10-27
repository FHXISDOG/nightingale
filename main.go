package main

import "github.com/gin-gonic/gin"

func main() {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messgae": "pong",
		})
	})
	//test msg
	r.GET("/test", func(c *gin.Context) {

	})

	r.Run()
}
