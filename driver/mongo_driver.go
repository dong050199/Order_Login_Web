package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	connStr := "mongodb://mongoAdmin:ttsxtm@128.199.124.231:27017/Becanatomy?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ok")
	Mongo.Client = client
	return Mongo

}
