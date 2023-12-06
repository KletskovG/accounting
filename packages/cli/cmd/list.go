package cmd

import (
	"encoding/json"
	"strconv"

	"github.com/cheynewallace/tabby"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/spf13/cobra"
)

func ListCommand(cmd *cobra.Command, args []string) {
	var limit int = 0

	if len(args) != 0 {
		if value, err := strconv.Atoi(args[0]); err == nil {
			limit = value
		}
	}

	var transactions = db.ListTransations(limit)
	logger.Info(len(transactions), " results")

	var results = tabby.New()
	results.AddHeader("Expense", "Date", "Category", "Note")

	for _, result := range transactions {
		results.AddLine(result.Expense, result.Date, result.Category, result.Note)
		transaction, err := json.MarshalIndent(result, "", "\t")

		if err != nil {
			logger.Error("Cant cast transaction to JSON, ", err, "\n", string(transaction))
		}

		if IsJsonFormat {
			logger.Info(string(transaction))
		}
	}

	if !IsJsonFormat {
		results.Print()
	}
}
