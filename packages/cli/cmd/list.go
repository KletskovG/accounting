package cmd

import (
	"strconv"

	"github.com/kletskovg/accounting/packages/db"
	"github.com/spf13/cobra"
)

func ListCommand(cmd *cobra.Command, args []string) {
	var limit int = 0

	if len(args) != 0 {
		if value, err := strconv.Atoi(args[0]); err == nil {
			limit = value
		}
	}

	db.ListTransations(limit)
}
