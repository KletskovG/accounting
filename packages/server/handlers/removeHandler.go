package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func deleteTransactions(response http.ResponseWriter, request *http.Request, docsToDelete int) {
	go db.RemoveTransaction(db.RemoveTransactionArgs{
		Count: docsToDelete,
	})
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "will delete "+strconv.Itoa(docsToDelete)+" transactions")
	return
}

func deleteTransactionByID(response http.ResponseWriter, request *http.Request, id string) {
	transactionID, idError := primitive.ObjectIDFromHex(id)

	if idError != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Cant process transaction id "+idError.Error())
		return
	}

	logger.Info("Transaction to be deleted", transactionID.Hex())

	go db.RemoveTransaction(db.RemoveTransactionArgs{
		IDs: []string{transactionID.Hex()},
	})

	logger.Info(transactionID.Hex(), "deleted")

	response.Header().Set(common.HeaderContentType, common.ContentTypeJson)
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "will delete transaction "+transactionID.Hex())
	return
}

func RemoveHandler(response http.ResponseWriter, request *http.Request) {
	if !request.URL.Query().Has("transaction") {
		go db.RemoveTransaction(db.RemoveTransactionArgs{
			RemoveLast: true,
		})

		response.Header().Set(common.HeaderContentType, common.ContentTypeJson)
		response.WriteHeader(http.StatusOK)
		// Maybe send transaction here
		io.WriteString(response, "Last transaction was deleted")
		return
	}

	transactionValue := request.URL.Query().Get("transaction")
	transactionsToDelete, err := strconv.Atoi(transactionValue)

	if err != nil {
		logger.Info(transactionValue)
		deleteTransactionByID(response, request, transactionValue)
		return
	}

	deleteTransactions(response, request, transactionsToDelete)
}
