package my_mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Collection struct {
	DatabaseName   string
	CollectionName string
}

var client *mongo.Client

func init() {
	client = GetClient()
}

func (c *Collection) GetCollection() *mongo.Collection {
	return client.Database(c.DatabaseName).Collection(c.CollectionName)
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func GetClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetMaxPoolSize(20))
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("mongo db connect error !!!")
	} else {
		fmt.Println("mongo db connect success!!!")
	}
	return client
}
