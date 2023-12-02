package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Command hugo was executed")
	},
}

func init() {
	RegisterCommand(&RegisterRecord{
		name: "add",
		root: rootCmd,
		impl: AddCommand,
	})
	RegisterCommand(&RegisterRecord{
		name: "remove",
		root: rootCmd,
		impl: RemoveCommand,
	})
	RegisterCommand(&RegisterRecord{
		name: "list",
		root: rootCmd,
		impl: ListCommand,
	})
	// RegisterReportCommand(rootCmd)
	// RegisterUpdateCommand(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
