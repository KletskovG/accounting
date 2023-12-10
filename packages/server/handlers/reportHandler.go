package handlers

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/accounting/server/services"
	"github.com/kletskovg/packages/common"
)

func isReportParamsEmpty(request *http.Request) bool {
	isStartPresented := request.URL.Query().Has("start")
	isEndPresented := request.URL.Query().Has("end")

	return !isStartPresented && !isEndPresented
}

func ReportHandler(response http.ResponseWriter, request *http.Request) {
	var start, end string

	if isReportParamsEmpty(request) {
		currentTimestamp := time.Now().UnixMilli()
		end = time.UnixMilli(currentTimestamp).Format(common.DateLayout)
		start = time.UnixMilli((currentTimestamp - int64(common.Month))).Format(common.DateLayout)
	} else if request.URL.Query().Has("days") {
		days, err := strconv.Atoi(request.URL.Query().Get("days"))

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			io.WriteString(response, "Provide start and end dates or number of days to process")
			return
		}

		currentTimestamp := time.Now().UnixMilli()
		end = time.UnixMilli(currentTimestamp).Format(common.DateLayout)
		start = time.UnixMilli((currentTimestamp - int64(common.Day)*int64(days))).Format(common.DateLayout)
	} else {
		start = request.URL.Query().Get("start")
		end = request.URL.Query().Get("end")
	}

	logger.Info("Dates", start, " ", end)

	transactions := db.ReportTransactions(start, end)
	report := common.GetCsvReport(transactions)
	response.WriteHeader(http.StatusOK)
	response.Header().Add(common.HeaderContentType, common.ContentTypeJson)
	io.WriteString(response, "OK, Wait for S3 link")

	go services.UploadReport(report, common.Hosts().TelegramAPIURL)
	return
}
