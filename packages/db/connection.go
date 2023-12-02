package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kletskovg/accounting/packages/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = new(mongo.Collection)

type TransactionsRepository struct{}

func init() {
	log.Default().Println("INIT OF DB WAS CALLED")
	var uri = config.GetEnvVariable(config.MONGODB_URL)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Cant connect to database")
		panic(err)
	}

	collection = client.Database(
		config.GetEnvVariable(config.MONGODB_NAME),
	).Collection(config.GetEnvVariable(config.MONGODB_COLLECTION))

	log.Default().Println("Collection: ", collection.Name(), " ready")
}

func Find() {
	title := "Back to the Future"
	var result bson.M
	var err = collection.FindOne(context.TODO(), bson.D{{"topups", 100}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
