package common

import (
	"strconv"
)

func GetCsvReport(transactions []Transaction) string {
	var csvReport = "Date,Amount,Category,Note\n"

	// TODO: Remove date duplicates
	for _, transaction := range transactions {
		expense := strconv.Itoa(int(transaction.Expense))

		csvLine := transaction.Date + "," + expense + "," + transaction.Category + "," + transaction.Note + "\n"
		csvReport += csvLine
	}

	return csvReport
}
