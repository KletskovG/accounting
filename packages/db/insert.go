package db

import (
	"context"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertTransaction(
	date string,
	expenseAmount int,
	category,
	note string,
) (*mongo.InsertOneResult, error) {
	transaction := Transaction{
		Date:      date,
		Expense:   int32(expenseAmount),
		Category:  category,
		Note:      note,
		Timestamp: time.Now().String(),
	}

	logger.Info("inserting: ", transaction)

	return Collection.InsertOne(context.Background(), transaction)
}
