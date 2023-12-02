package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Work with limit
func ListTransations(limit int) {
	var opts = options.Find().SetLimit(int64(limit))
	var cursor, err = Collection.Find(context.Background(), bson.D{}, opts)

	if err != nil {
	}

	// var transactions = make(map[string]bson.M)

	defer cursor.Close(context.Background())

	var transactions []Transaction

	if err = cursor.All(context.Background(), &transactions); err != nil {
		log.Default().Println("Cant read transactions from DB: ", err)
	}

	log.Default().Println(len(transactions), " results")
	for _, result := range transactions {
		// log.Default().Println(result.Timestamp.T)
		log.Default().Println(result)
	}
}
