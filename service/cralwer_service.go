package service

import (
	"mycode/nightingale/crawler"
	"mycode/nightingale/my_mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserInfoMap map[string]*UserInfo = make(map[string]*UserInfo)

type UserInfo struct {
	MsgChan  chan *crawler.ParseResult
	flagChan chan int
	ChanLen  int
}

func GetMsg(c *gin.Context) interface{} {
	result := make([]crawler.ParseResult, 0, 3)
	// 	id := c.Query("id")
	// 	refreshFlag := c.DefaultQuery("refreshFlag", "0")
	// 	if refreshFlag == "1" {
	// 		delete(UserInfoMap, id)
	// 	}
	// 	userInfo, ok := UserInfoMap[id]
	// 	if !ok {
	// 		rules := reader.ReadRss("/Users/finger/code/mycode/nightingale/rss.json")
	// 		userInfo = &UserInfo{
	// 			MsgChan:  make(chan *crawler.ParseResult, 3),
	// 			flagChan: make(chan int, len(rules)),
	// 			ChanLen:  len(rules),
	// 		}
	// 		for _, val := range rules {
	// 			go func(r crawler.XmlRule) {
	// 				ch := r.GenerateMsgChan()
	// 				for {
	// 					as, ok := <-ch
	// 					if !ok {
	// 						userInfo.flagChan <- 1
	// 						break
	// 					}
	// 					userInfo.MsgChan <- as
	// 				}
	// 			}(val)
	// 		}
	// 		UserInfoMap[id] = userInfo
	// 	}
	// 	for {
	// 		select {
	// 		case <-userInfo.flagChan:
	// 			userInfo.ChanLen--
	// 			if userInfo.ChanLen <= 0 {
	// 				userInfo.flagChan = nil
	// 				close(userInfo.MsgChan)
	// 			}
	// 		case val, ok := <-userInfo.MsgChan:
	// 			if !ok {
	// 				goto END
	// 			}
	// 			result = append(result, *val)
	// 			if len(result) == 3 {
	// 				goto END
	// 			}
	// 		}
	// 	}
	// END:
	return result
}

func RuleFindAll(c *gin.Context) interface{} {
	return my_mongo.GetAllRule()
}

func FindAllRule(c *gin.Context) interface{} {
	return my_mongo.GetAllRule()
}

func FindRuleMsg(c *gin.Context) interface{} {
	id := c.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	rule := my_mongo.FindOne(bson.M{"_id": objectID})
	return rule.GetResult()
}
