package handlers

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/packages/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddHandler(response http.ResponseWriter, request *http.Request) {
	amountValue := request.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad amount")
		return
	}

	dateValue := request.URL.Query().Get("date")
	date, err := time.Parse(common.DateLayout, dateValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad date")
		return
	}

	category := request.URL.Query().Get("category")
	note := request.URL.Query().Get("note")

	transaction, err := db.InsertTransaction(date.Format(common.DateLayout), amount, category, note)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, "Cant write transaction to DB")
		return
	}

	transactionID, ok := transaction.InsertedID.(primitive.ObjectID)
	if !ok {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, "Cant write transaction to DB")
		return
	}

	response.Header().Set(common.HeaderContentType, common.ContentTypeJson)
	response.WriteHeader(http.StatusCreated)
	io.WriteString(response, transactionID.String())
}
