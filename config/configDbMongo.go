package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GlobalDbMongo client
var GlobalDbMongo *mongo.Client
var ctx = context.Background()

//InitMongo client
func InitMongo(mongoConfig MongoDatabase) (client *mongo.Client, err error) {
	clientOptions := options.Client().ApplyURI(mongoConfig.Connection).SetRetryWrites(false)

	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, err
}

func init() {
	db := MongoDatabase{
		Connection: "mongodb://127.0.0.1:27017",
		Database:   "kygo",
	}
	GlobalDbMongo, _ = InitMongo(db)
}
