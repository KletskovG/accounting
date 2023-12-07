package db

import (
	"context"
	"sync"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func ReportTransactions(start, end string) {
	startDate, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)

	startTimestamp := startDate.UTC().UnixMilli()
	endTimestamp := endDate.UTC().UnixMilli()
	cursor, err := Collection.Find(context.Background(), bson.D{})

	if err != nil {
		logger.Info("Can't find transactions: ", err)
		return
	}

	defer cursor.Close(context.Background())

	var transactions []Transaction
	if err := cursor.All(context.Background(), &transactions); err != nil {
		logger.Info("Cant process transactions cursor, ", err)
		return
	}

	var wg sync.WaitGroup
	var results []Transaction = make([]Transaction, 0)

	for _, result := range transactions {
		wg.Add(1)
		go func(result Transaction) {
			defer wg.Done()
			transactionTimestamp, err := time.Parse("2006-01-02", result.Date)

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
	logger.Info("RESULTS")
	logger.Info(results)
}
