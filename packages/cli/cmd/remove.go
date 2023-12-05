package cmd

import (
	"strconv"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/spf13/cobra"
)

func RemoveCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		db.RemoveTransaction(db.RemoveTransactionArgs{
			RemoveLast: true,
		})
		return
	}

	docsToDelete, err := strconv.Atoi(args[0])

	if err == nil {
		db.RemoveTransaction(db.RemoveTransactionArgs{
			Count: docsToDelete,
		})
		return
	}

	db.RemoveTransaction(db.RemoveTransactionArgs{
		IDs: args,
	})
}
