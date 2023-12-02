package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kletskovg/accounting/packages/cli/utils"
	"github.com/kletskovg/accounting/packages/db"
	"github.com/spf13/cobra"
)

const ADD_REQUIRED_ARGS = 3

type AddArgs struct {
	date          string
	expenseAmount int
	category      string
	note          string
}

func AddCommand(cmd *cobra.Command, args []string) {
	if len(args) < ADD_REQUIRED_ARGS {
		fmt.Println("usage: add <date> <amount> <category?> <note?>")
		return
	}

	var amount, error = strconv.Atoi(args[1])

	if error != nil {
		fmt.Println(error)
		return
	}

	// TODO: seems struct is not needed here
	addArgs := &AddArgs{
		date:          args[0],
		expenseAmount: amount,
		category:      utils.ReadArgByIndex(args, 2),
		note:          utils.ReadArgByIndex(args, 3),
	}

	var result, err = db.InsertTransaction(addArgs.date, addArgs.expenseAmount, addArgs.category, addArgs.note)

	if err != nil {
		log.Default().Println("ERROR: Cant insert transaction \n", err)
	}

	log.Default().Println("Transaction added: \n", &result.InsertedID)
}
