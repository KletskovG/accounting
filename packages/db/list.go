package db

import (
	"context"

	"github.com/kletskovg/accounting/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListTransations(limit int) []primitive.D {
	var opts = options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{{"timestamp", 1}})

	var cursor, err = Collection.Find(
		context.Background(),
		bson.D{},
		opts,
	)

	if err != nil {
		logger.Error("Cant retrieve transactions from DB \n")
	}

	defer cursor.Close(context.Background())

	var transactions []bson.D

	if err = cursor.All(context.Background(), &transactions); err != nil {
		logger.Error("Cant process transactions from DB: \n", err)
	}

	return transactions
}
