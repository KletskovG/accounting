package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RemoveCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Should remove last transation")
		return
	}

	var transactionId = args[0]
	fmt.Println("ID of transation which should be removed: ", transactionId)
}
