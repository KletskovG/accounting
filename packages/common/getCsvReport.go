package common

import (
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/kletskovg/accounting/packages/logger"
)

func mergeTransactionsByCategory(transactionsInDay []Transaction) []Transaction {
	transactions := make(map[string]Transaction, 0)

	for _, transaction := range transactionsInDay {
		categoryName := strings.ToLower(transaction.Category)
		transactionNote := strings.ToLower(transaction.Note)
		transactionKey := categoryName + "-" + transactionNote
		if acc, ok := transactions[categoryName]; ok {
			transactions[transactionKey] = Transaction{
				Expense:  int32(acc.Expense) + int32(transaction.Expense),
				Category: acc.Category,
				Date:     acc.Date,
				Note:     acc.Note + " " + transaction.Note,
			}
		} else {
			transactions[transactionKey] = transaction
		}
	}

	result := []Transaction{}

	for _, value := range transactions {
		result = append(result, value)
	}

	return result
}

func splitTransactionsByDays(transactions []Transaction) [][]Transaction {
	result := [][]Transaction{}
	transactionsBatch := []Transaction{}
	startDate, _ := time.Parse(DateLayout, transactions[0].Date)
	startTimestamp := startDate.UnixMilli()

	for _, transaction := range transactions {
		currentDate, _ := time.Parse(DateLayout, transaction.Date)
		currentTimestamp := currentDate.UnixMilli()

		if currentTimestamp == startTimestamp {
			transactionsBatch = append(transactionsBatch, transaction)
		} else {
			result = append(result, transactionsBatch)
			transactionsBatch = []Transaction{transaction}
			startTimestamp = currentTimestamp
		}
	}

	result = append(result, transactionsBatch)

	return result
}

func fillEmptyDays(startDate, endDate time.Time) string {
	result := []string{}
	logger.Info("DATES")
	logger.Info(startDate, endDate)
	start := startDate
	daysBetween := int(math.RoundToEven(endDate.Sub(startDate).Hours() / 24))
	for i := 0; i < daysBetween; i++ {
		start = start.Add(time.Hour * 24)
		result = append(result, start.Format(DateLayout)+",,,\n")
	}

	slices.Reverse(result)

	return ReduceSlice(result, func(dates, current string) string {
		dates += current
		return dates
	}, "")
}

func GetCsvReport(transactions []Transaction) string {
	var csvReport = "Date,Amount,Category,Note\n"
	transactionsReversed := slices.Clone(transactions)
	slices.Reverse(transactionsReversed)

	days := splitTransactionsByDays(transactionsReversed)

	for index, day := range days {
		days[index] = mergeTransactionsByCategory(day)
	}

	amountOfDays := len(days)
	for dayIndex, day := range days {
		currentDate, _ := time.Parse(DateLayout, day[0].Date)

		for _, transaction := range day {
			expense := strconv.Itoa(int(transaction.Expense))

			csvLine := transaction.Date + "," + expense + "," + transaction.Category + "," + transaction.Note + "\n"
			csvReport += csvLine
		}

		if dayIndex < amountOfDays-1 {
			logger.Info(dayIndex+1, amountOfDays)
			prevDayDate, _ := time.Parse(DateLayout, days[dayIndex+1][0].Date)
			dateDiff := currentDate.Sub(prevDayDate)

			if dateDiff.Hours() > 24 {
				csvReport += fillEmptyDays(prevDayDate, currentDate)
			}
		}

	}

	return csvReport
}
