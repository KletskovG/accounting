package db

import (
	"context"
	"log"

	"github.com/kletskovg/accounting/packages/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection = new(mongo.Collection)

func init() {
	var uri = config.GetEnvVariable(config.MONGODB_URL)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Cant connect to database")
		panic(err)
	}

	Collection = client.Database(
		config.GetEnvVariable(config.MONGODB_NAME),
	).Collection(config.GetEnvVariable(config.MONGODB_COLLECTION))
}
