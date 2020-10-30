package my_mongo

import (
	"context"
	"mycode/nightingale/base"
	"mycode/nightingale/crawler"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRule struct {
	crawler.Rule
}

var database *Database

var collection *Collection

func init() {
	database = &Database{DatabaseName: "nightingale"}
	collection = &Collection{Database: database, CollectionName: "rule"}
}

func (rule *MongoRule) FindAll() base.Any {
	var result crawler.XmlRule
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("nightingale").Collection("rule")
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	sr := collection.FindOne(ctx, bson.M{"DateNode": "//pubDate"})
	sr.Decode(&result)
	return result
}

func (rule *MongoRule) FindOne() base.Any {
	return nil
}
