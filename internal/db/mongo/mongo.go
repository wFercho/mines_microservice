package db

import (
	"context"
	"log"

	"github.com/wFercho/mines_microservice/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MONGO_DB *mongo.Database

func ConnectToMongoDatabase() {
	cfg := config.LoadConfig()
	uri := cfg.GetMongoDatabaseURL()
	docs := "www.mongodb.com/docs/drivers/go/current/"

	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	MONGO_DB = client.Database(cfg.MongoDB.DBName)
}
