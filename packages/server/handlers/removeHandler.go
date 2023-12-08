package handlers

import (
	"io"
	"net/http"

	"github.com/kletskovg/accounting/packages/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RemoveHandler(response http.ResponseWriter, request *http.Request) {
	transactionIDValue := request.URL.Query().Get("id")
	transactionID, err := primitive.ObjectIDFromHex(transactionIDValue)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		io.WriteString(response, "Bad transaction id")
		return
	}

	// TODO: Stopped here
	// Need to support all remove arguments
	err = db.RemoveTransaction(&db.RemoveTransactionArgs{
		transactionID
	})
}
