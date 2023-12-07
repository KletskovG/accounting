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

const (
	addRequiredArgs  = 1
	amountArgIndex   = 0
	categoryArgIndex = 1
	dateArgIndex     = 2
	noteArgIndex     = 3
)

type AddArgs struct {
	date          string
	expenseAmount int
	category      string
	note          string
}

func AddArgsValidator(cmd *cobra.Command, args []string) error {
	if len(args) >= addRequiredArgs {
		return nil
	}
	var argError = errors.New("add command requires minimum 1 argument - amount of expense")
	logger.Info(argError)
	return argError
}

func AddCommand(cmd *cobra.Command, args []string) {
	var amount, error = strconv.Atoi(args[amountArgIndex])

	if error != nil {
		logger.Error(error)
		return
	}

	var category = utils.ReadArgByIndex(args, categoryArgIndex)

	if category == "" {
		category = "Other"
	}

	var date = utils.ParseDate(
		utils.ReadArgByIndex(args, dateArgIndex),
	)

	if date == "" {
		date = time.Now().UTC().Format(common.DateLayout)
	}

	addArgs := &AddArgs{
		expenseAmount: amount,
		date:          date,
		category:      category,
		note:          utils.ReadArgByIndex(args, noteArgIndex),
	}

	logger.Info("Inserting transaction with args: \n", addArgs)

	var result, err = db.InsertTransaction(addArgs.date, addArgs.expenseAmount, addArgs.category, addArgs.note)

	if err != nil {
		logger.Error("Cant insert transaction \n", err)
	}

	logger.Info("Transaction added: \n", &result.InsertedID)
}
