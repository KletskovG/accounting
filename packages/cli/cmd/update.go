package cmd

import (
	"errors"
	"strconv"

	"github.com/kletskovg/accounting/packages/cli/utils"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"github.com/spf13/cobra"
)

const (
	updateMinArgs       = 2
	transactionArgIndex = 0
)

func UpdateArgsValidator(cmd *cobra.Command, args []string) error {
	if len(args) >= updateMinArgs {
		return nil
	}

	var argError = errors.New("update command requires 2 arguments minimum - ID of transaction and amount")
	logger.Info(argError)
	return argError
}

func UpdateCommand(cmd *cobra.Command, args []string) {
	var transactionID = args[0]
	var amount, err = strconv.Atoi(args[1])
	var category = utils.ReadArgByIndex(args, 2)
	var date = utils.ReadArgByIndex(args, 3)
	var note = utils.ReadArgByIndex(args, 4)
	if err != nil {
		logger.Error("Cant convert expense to number, ", err)
	}

	var result = db.UpdateTransaction(transactionID, &common.Transaction{
		Expense:  int32(amount),
		Note:     note,
		Date:     date,
		Category: category,
	})

	if result.Err() != nil {
		logger.Error("Cant update transaction in DB, ", result.Err().Error())
	}

	logger.Info("Transaction update, ", result)
}
