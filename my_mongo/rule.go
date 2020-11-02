package my_mongo

import (
	"context"
	"mycode/nightingale/crawler"

	"go.mongodb.org/mongo-driver/bson"
)

var collection *Collection

func init() {
	collection = &Collection{"nightingale", "rule"}
}

func GetAllRule() interface{} {
	var result []crawler.Rule
	coll := collection.GetCollection()
	sr, _ := coll.Find(context.TODO(), bson.M{})
	sr.All(context.TODO(), &result)
	return result
}

func FindOne(cond map[string]interface{}) crawler.Rule {
	var result crawler.Rule
	collection.GetCollection().FindOne(context.TODO(), cond).Decode(&result)
	return result
}
