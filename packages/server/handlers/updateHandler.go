package handlers

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
)

func UpdateHandler(response http.ResponseWriter, request *http.Request) {
	if !request.URL.Query().Has("transaction") {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Need transaction id to update")
		return
	}

	transactionValue := request.URL.Query().Get("transaction")

	amountValue := request.URL.Query().Get("amount")
	logger.Info("Amount,", amountValue)
	amount, err := strconv.Atoi(amountValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad amount")
		return
	}

	date, err := time.Parse(common.DateLayout, request.URL.Query().Get("date"))
	var dateValue string

	if err != nil {
		dateValue = ""
	}

	dateValue = date.String()

	category := request.URL.Query().Get("category")
	note := request.URL.Query().Get("note")

	updateResult := db.UpdateTransaction(transactionValue, &common.Transaction{
		Expense:  int32(amount),
		Note:     note,
		Date:     dateValue,
		Category: category,
	})

	if updateResult.Err() != nil {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, "Cant update transaction"+updateResult.Err().Error())
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Header().Add(common.HeaderContentType, common.ContentTypeJson)
	io.WriteString(response, string(transactionValue))
}
