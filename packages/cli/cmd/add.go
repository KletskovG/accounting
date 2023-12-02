package cmd

import (
	"fmt"
	"strconv"

	"github.com/kletskovg/accounting/packages/cli/utils"
	"github.com/spf13/cobra"
)

var addRequiredArgs = 3

type AddArgs struct {
	date          string
	expenseAmount int
	category      string
	note          string
}

func AddCommand(cmd *cobra.Command, args []string) {
	if len(args) < addRequiredArgs {
		fmt.Println("usage: add <date> <amount> <category?> <note?>")
		return
	}

	var amount, error = strconv.Atoi(args[1])

	if error != nil {
		fmt.Println(error)
		return
	}

	addArgs := &AddArgs{
		date:          args[0],
		expenseAmount: amount,
		category:      utils.ReadArgByIndex(args, 2),
		note:          utils.ReadArgByIndex(args, 3),
	}

	// TODO: write transation to DB here
	fmt.Println("Args: ", addArgs)
}
