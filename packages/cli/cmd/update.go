package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RegisterUpdateCommand(rootCmd *cobra.Command) {
	var UpdateCommand = &cobra.Command{
		Use:   "update",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Update command was executed")
		},
	}

	rootCmd.AddCommand(UpdateCommand)
}
