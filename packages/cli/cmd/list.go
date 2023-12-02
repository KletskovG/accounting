package cmd

import (
	"strconv"

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
	for _, result := range transactions {
		logger.Info(result)
	}

	// TODO: Implement indentation and logging of transactions with ID
}
