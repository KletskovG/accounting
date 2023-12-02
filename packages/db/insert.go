package db

import (
	"context"
	"log"
	"time"

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

	log.Default().Println("inserting: ", transaction)

	return Collection.InsertOne(context.Background(), transaction)
}
