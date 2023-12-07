package cmd

import (
	"errors"

	"github.com/kletskovg/accounting/packages/cli/utils"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/spf13/cobra"
)

func ReportArgsValidator(rootCmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		err := errors.New("report requires 2 args - start date and end date")
		logger.Info(err)
		return err
	}

	var start = utils.ParseDate(
		utils.ReadArgByIndex(args, 0),
	)
	var end = utils.ParseDate(
		utils.ReadArgByIndex(args, 1),
	)

	if start == "" || end == "" {
		err := errors.New("Start end empty is required")
		logger.Info(err)
		return err
	}

	return nil
}

// last month - default
// Option - day since now
// Option - range of dates

func ReportCommand(rootCmd *cobra.Command, args []string) {
	start := utils.ParseDate(
		utils.ReadArgByIndex(args, 0),
	)
	end := utils.ParseDate(
		utils.ReadArgByIndex(args, 1),
	)

	// TODO: Finished here
	db.ReportTransactions(start, end)
}
