package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kletskovg/accounting/packages/cli/utils"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/spf13/cobra"
)

const AddRequiredArgs = 1

type AddArgs struct {
	date          string
	expenseAmount int
	category      string
	note          string
}

func AddCommand(cmd *cobra.Command, args []string) {
	if len(args) < AddRequiredArgs {
		fmt.Println("usage: add   <amount> <date?> <category?> <note?>")
		return
	}

	var amount, error = strconv.Atoi(args[0])

	if error != nil {
		fmt.Println(error)
		return
	}

	var date = utils.ReadArgByIndex(args, 1)

	if date == "" {
		date = time.Now().UTC().String()
	}

	var category = utils.ReadArgByIndex(args, 2)

	if category == "" {
		category = "Other"
	}

	// TODO: seems struct is not needed here
	addArgs := &AddArgs{
		expenseAmount: amount,
		date:          date,
		category:      category,
		note:          utils.ReadArgByIndex(args, 3),
	}

	logger.Info("Inserting transaction with args: \n", addArgs)

	var result, err = db.InsertTransaction(addArgs.date, addArgs.expenseAmount, addArgs.category, addArgs.note)

	if err != nil {
		logger.Error("Cant insert transaction \n", err)
	}

	logger.Info("Transaction added: \n", &result.InsertedID)
}
