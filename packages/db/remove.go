package db

import (
	"context"
	"sync"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RemoveTransactionArgs struct {
	RemoveLast bool
	Count      int
	IDs        []string
}

func RemoveTransaction(args RemoveTransactionArgs) {
	if args.RemoveLast {
		logger.Info("Try to remove last transaction")
		var findOpts = options.FindOneAndDeleteOptions{
			Sort: bson.D{{Key: "timestamp", Value: -1}},
		}
		Collection.FindOneAndDelete(context.Background(), bson.D{}, &findOpts)
		return
	}

	var wg sync.WaitGroup

	if args.Count > 0 {
		logger.Info("Try to remove ", args.Count, " transactions")

		var findOpts = options.Find()
		findOpts.SetLimit(int64(args.Count))
		findOpts.SetSort(bson.D{{Key: "timestamp", Value: -1}})

		var cursor, err = Collection.Find(context.Background(), bson.D{}, findOpts)

		if err != nil {
			logger.Error("Cant find documents from DB", err)
		}

		var transactions []common.Transaction

		if err = cursor.All(context.Background(), &transactions); err != nil {
			logger.Error("Cant process transactions from DB: \n", err)
		}

		for _, transaction := range transactions {
			wg.Add(1)

			go func(deleteTransaction common.Transaction) {
				defer wg.Done()
				logger.Info("Deleting...", deleteTransaction)
				var transactionID, err = primitive.ObjectIDFromHex(deleteTransaction.ID)

				if err == nil {
					Collection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: transactionID}})
				} else {
					logger.Error("Cant create ObjectID, ", err, deleteTransaction)
				}
			}(transaction)
		}
		wg.Wait()

		return
	}

	if len(args.IDs) > 0 {
		for _, id := range args.IDs {
			wg.Add(1)
			var objectID, err = primitive.ObjectIDFromHex(id)

			if err != nil {
				logger.Info("Cant process ", id, "not a valid object id")
				continue
			}

			go func(docId primitive.ObjectID) {
				defer wg.Done()
				Collection.FindOneAndDelete(
					context.Background(),
					bson.D{{
						Key: "_id", Value: docId,
					}},
				)
			}(objectID)
		}
		wg.Wait()
	}
}
