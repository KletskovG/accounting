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

func UpdateHandler(response http.ResponseWriter, request *http.Request) {
	if !request.URL.Query().Has("transaction") {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Need transaction id to update")
		return
	}

	transactionValue := request.URL.Query().Get("transaction")
	transactionID, idError := primitive.ObjectIDFromHex(transactionValue)

	if idError != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Cant process transaction id "+idError.Error())
		return
	}

	amountValue := request.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad amount")
		return
	}

	date, err := time.Parse(common.DateLayout, request.URL.Query().Get("date"))
	var dateValue string

	if err != nil {
		// TODO: Update default date value - right now it breaks logic, ovverrides current date
		dateValue = ""
	}

	dateValue = date.String()

	category := request.URL.Query().Get("category")
	note := request.URL.Query().Get("note")

	updateResult := db.UpdateTransaction(transactionID.Hex(), &common.Transaction{
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
	io.WriteString(response, string(transactionID.Hex()))
}
