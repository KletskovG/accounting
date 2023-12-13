package db

import (
	"context"
	"sort"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson"
)

func ReportTransactions(start, end string) []common.Transaction {
	startDate, _ := time.Parse(common.DateLayout, start)
	endDate, _ := time.Parse(common.DateLayout, end)

	logger.Info(startDate, endDate)

	startTimestamp := startDate.UTC().UnixMilli()
	endTimestamp := endDate.UTC().UnixMilli()
	cursor, err := Collection.Find(context.Background(), bson.D{})

	if err != nil {
		logger.Info("Can't find transactions: ", err)
		return []common.Transaction{}
	}

	defer cursor.Close(context.Background())

	var transactions []common.Transaction
	if err := cursor.All(context.Background(), &transactions); err != nil {
		logger.Info("Cant process transactions cursor, ", err)
		return []common.Transaction{}
	}

	var results []common.Transaction = make([]common.Transaction, 0)

	for _, result := range transactions {
		transactionTimestamp, err := time.Parse(common.DateLayout, result.Date)

		if err != nil {
			logger.Info("Cant parse date", result.Date, err)
			continue
		}

		if transactionTimestamp.UTC().UnixMilli() >= startTimestamp && transactionTimestamp.UTC().UnixMilli() <= endTimestamp {
			results = append(results, result)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		leftTimestamp, startErr := time.Parse(common.DateLayout, results[i].Date)
		rightTimestamp, endErr := time.Parse(common.DateLayout, results[j].Date)

		if startErr != nil || endErr != nil {
			logger.Info("Report: cant parse dates to sort transactions, ", startErr, endErr)
			return false
		}

		return leftTimestamp.UnixMilli() > rightTimestamp.UnixMilli()
	})

	return results
}
