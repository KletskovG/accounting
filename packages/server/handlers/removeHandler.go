package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
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

	logger.Info("Transaction to be deleted", transactionID.String())

	go db.RemoveTransaction(db.RemoveTransactionArgs{
		IDs: []string{transactionID.String()},
	})

	logger.Info(transactionID.String(), "deleted")

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "will delete transaction "+transactionID.String())
	return
}

// TODO: Support remove last option
// TODO: Stopped here, somehow cant delete by id, also need to delete last transaction
func RemoveHandler(response http.ResponseWriter, request *http.Request) {
	transactionValue := request.URL.Query().Get("transaction")
	transactionsToDelete, err := strconv.Atoi(transactionValue)

	if err != nil {
		logger.Info(transactionValue)
		deleteTransactionByID(response, request, transactionValue)
		return
	}

	deleteTransactions(response, request, transactionsToDelete)
}
