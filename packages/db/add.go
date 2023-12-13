package db

import (
	"context"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTransaction(
	date string,
	expenseAmount int,
	category,
	note string,
) (transactionID string, insertionError error) {
	transaction := common.Transaction{
		Date:      date,
		Expense:   int32(expenseAmount),
		Category:  category,
		Note:      note,
		Timestamp: time.Now().String(),
	}

	logger.Info("inserting: ", transaction)

	result, err := Collection.InsertOne(context.Background(), transaction)

	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).String(), nil
}
