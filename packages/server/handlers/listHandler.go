package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
)

type ListHandler struct{}

func (lh ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	limitValue := request.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitValue)

	if err != nil {
		limit = 0
	}

	transactions := db.ListTransations(limit)

	response.Header().Set(common.HeaderContentType, common.ContentTypeJson)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(transactions)

	if err != nil {
		logger.Info("Cant cast transactions to JSON, ", err)
	}
	response.WriteHeader(http.StatusOK)
	io.WriteString(response, string(result))
}
