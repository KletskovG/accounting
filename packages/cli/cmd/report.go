package cmd

import (
	"errors"
	"strconv"
	"time"

	"github.com/kletskovg/accounting/packages/cli/utils"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
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

// Option - day since now

func ReportCommand(rootCmd *cobra.Command, args []string) {
	var start, end string

	if len(args) == 0 {
		currentTimestamp := time.Now().UnixMilli()
		end = time.UnixMilli(currentTimestamp).Format(common.DateLayout)
		start = time.UnixMilli((currentTimestamp - int64(common.Month))).Format(common.DateLayout)
	} else if len(args) == 1 {
		days, err := strconv.Atoi(utils.ReadArgByIndex(args, 0))

		if err != nil {
			logger.Info("Provide valid amount of days", err)
			return
		}

		currentTimestamp := time.Now().UnixMilli()
		end = time.UnixMilli(currentTimestamp).Format(common.DateLayout)
		start = time.UnixMilli((currentTimestamp - int64(common.Day)*int64(days))).Format(common.DateLayout)
	} else {
		start = utils.ParseDate(
			utils.ReadArgByIndex(args, 0),
		)
		end = utils.ParseDate(
			utils.ReadArgByIndex(args, 1),
		)
	}

	// TODO: Finished here
	db.ReportTransactions(start, end)
}
