package db

import (
	"context"
	"sync"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson"
)

func ReportTransactions(start, end string) []Transaction {
	startDate, _ := time.Parse(common.DateLayout, start)
	endDate, _ := time.Parse(common.DateLayout, end)

	logger.Info(startDate, endDate)

	startTimestamp := startDate.UTC().UnixMilli()
	endTimestamp := endDate.UTC().UnixMilli()
	cursor, err := Collection.Find(context.Background(), bson.D{})

	if err != nil {
		logger.Info("Can't find transactions: ", err)
		return []Transaction{}
	}

	defer cursor.Close(context.Background())

	var transactions []Transaction
	if err := cursor.All(context.Background(), &transactions); err != nil {
		logger.Info("Cant process transactions cursor, ", err)
		return []Transaction{}
	}

	var wg sync.WaitGroup
	var results []Transaction = make([]Transaction, 0)

	for _, result := range transactions {
		wg.Add(1)
		go func(result Transaction) {
			defer wg.Done()
			transactionTimestamp, err := time.Parse(common.DateLayout, result.Date)

			if err != nil {
				logger.Info("Cant parse date, ", result.Date, err)
				return
			}

			if transactionTimestamp.UTC().UnixMilli() >= startTimestamp && transactionTimestamp.UTC().UnixMilli() <= endTimestamp {
				results = append(results, result)
			}
		}(result)
	}

	wg.Wait()

	return results
}
