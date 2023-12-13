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

type AddHandler struct{}

func (ah AddHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	amountValue := request.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad amount")
		return
	}

	dateValue := request.URL.Query().Get("date")
	logger.Info("DATE", dateValue)
	logger.Info(amount)
	date, err := time.Parse(common.DateLayout, dateValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad date")
		return
	}

	category := request.URL.Query().Get("category")
	note := request.URL.Query().Get("note")

	transactionID, err := db.InsertTransaction(date.Format(common.DateLayout), amount, category, note)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, "Cant write transaction to DB"+err.Error())
	}

	response.WriteHeader(http.StatusCreated)
	response.Header().Set(common.HeaderContentType, common.ContentTypeJson)
	io.WriteString(response, transactionID)
}
