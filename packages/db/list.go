package db

import (
	"context"
	"sort"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListTransations(limit int) []common.Transaction {
	var opts = options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.D{{"timestamp", -1}})

	var cursor, err = Collection.Find(
		context.Background(),
		bson.D{},
		opts,
	)

	if err != nil {
		logger.Error("Cant retrieve transactions from DB \n")
	}

	defer cursor.Close(context.Background())

	var transactions []common.Transaction

	if err = cursor.All(context.Background(), &transactions); err != nil {
		logger.Error("Cant process transactions from DB: \n", err)
	}

	// TODO: Remove duplicates from report.go
	sort.Slice(transactions, func(i, j int) bool {
		leftTimestamp, startErr := time.Parse(common.DateLayout, transactions[i].Date)
		rightTimestamp, endErr := time.Parse(common.DateLayout, transactions[j].Date)

		if startErr != nil || endErr != nil {
			logger.Info("Report: cant parse dates to sort transactions, ", startErr, endErr)
			return false
		}

		return leftTimestamp.UnixMilli() > rightTimestamp.UnixMilli()
	})

	return transactions
}
