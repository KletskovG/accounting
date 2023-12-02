package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type RegisterRecord struct {
	name string
	root *cobra.Command
	impl func(cmd *cobra.Command, args []string)
}

func RegisterCommand(record *RegisterRecord) {
	var command = &cobra.Command{
		Use: record.name,
		Run: record.impl,
		PreRun: func(_ *cobra.Command, _ []string) {
			var cmd = exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		},
	}

	record.root.AddCommand(command)
}
