package service

import (
	"fmt"
	"mycode/nightingale/crawler"
	"mycode/nightingale/reader"

	"github.com/gin-gonic/gin"
)

var chanMap map[string]chan *crawler.ParseResult = make(map[string]chan *crawler.ParseResult)

func GetMsg(c *gin.Context) {
	result := make([]crawler.ParseResult, 0, 3)
	id := c.Query("id")
	fmt.Println(id)
	channel, ok := chanMap[id]
	if !ok {
		rule := reader.ReadRss("/Users/finger/code/mycode/nightingale/rss.json")
		channel = make(chan *crawler.ParseResult, 3)
		for i := range rule {
			go rule[i].GenerateResult(channel)
		}
		chanMap[id] = channel
	}
	for val := range channel {
		result = append(result, (*val))
		fmt.Println(len(result))
		if len(result) == 3 {
			break
		}
	}
	c.JSON(200, gin.H{
		"result": result,
	})
}

func closeChannel(c *gin.Context) {

}
