package db

import (
	"context"
	"log"

	"github.com/kletskovg/accounting/packages/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection = new(mongo.Collection)

type Transaction struct {
	ID        string `bson:"_id" json:"_id"`
	Date      string `bson:"date" json:"date"`
	Expense   int32  `bson:"expense" json:"expense"`
	Category  string `bson:"category" json:"category"`
	Note      string `bson:"note" json:"note"`
	Timestamp string `bson:"timestamp" json:"timestamp"`
}

func init() {
	var uri = config.GetEnvVariable(config.MONGODB_URL)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Cant connect to database")
		panic(err)
	}

	Collection = client.Database(
		config.GetEnvVariable(config.MONGODB_NAME),
	).Collection(config.GetEnvVariable(config.MONGODB_COLLECTION))
}
