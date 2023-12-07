package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type RegisterRecord struct {
	name          string
	root          *cobra.Command
	impl          func(cmd *cobra.Command, args []string)
	argsValidator func(cmd *cobra.Command, args []string) error
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
		Args: func(cmd *cobra.Command, args []string) error {
			if record.argsValidator == nil {
				return nil
			}

			return record.argsValidator(cmd, args)
		},
	}

	record.root.AddCommand(command)
}
