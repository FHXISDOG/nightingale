package my_mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	DatabaseName string
}

type Collection struct {
	*Database
	CollectionName string
}

var client *mongo.Client
var Ctx context.Context

func init() {
	client, Ctx = getConnection()
}

func (c *Collection) GetCollection() *mongo.Collection {
	return client.Database(c.DatabaseName).Collection(c.CollectionName)
}

func (d *Database) GetDatabase() *mongo.Database {
	return client.Database(d.DatabaseName)
}

func (d *Database) GetCollection(collectionName string) *mongo.Collection {
	return client.Database(d.DatabaseName).Collection(collectionName)
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func GetDatabase(databaseName string) *mongo.Database {
	return client.Database(databaseName)
}

func getConnection() (*mongo.Client, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("mongo db connect error !!!")
	} else {
		fmt.Println("mongo db connect success!!!")
	}
	return client, ctx
}
