package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
)

func deleteTransactions(response http.ResponseWriter, request *http.Request, docsToDelete int) {
	go db.RemoveTransaction(db.RemoveTransactionArgs{
		Count: docsToDelete,
	})
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "will delete "+strconv.Itoa(docsToDelete)+" transactions")
}

func deleteTransactionByID(response http.ResponseWriter, request *http.Request, id string) {
	logger.Info("Transaction to be deleted", id)

	go db.RemoveTransaction(db.RemoveTransactionArgs{
		IDs: []string{id},
	})

	logger.Info(id, "deleted")

	response.Header().Set(common.HeaderContentType, common.ContentTypeJson)
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, "will delete transaction "+id)
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
