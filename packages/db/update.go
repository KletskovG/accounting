package db

import (
	"context"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTransaction(
	id string,
	// expenseAmount int,
	// date,
	// category,
	// note string,
	transaction *Transaction,
) *mongo.SingleResult {
	transactionID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("Cant get ID of transaction, ", err, id)
	}

	transaction.Timestamp = time.Now().String()

	return Collection.FindOneAndUpdate(
		context.Background(),
		bson.D{{Key: "_id", Value: transactionID}},
		bson.D{{Key: "$set", Value: transaction}},
	)
}
