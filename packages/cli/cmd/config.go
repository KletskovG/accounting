package cmd

import (
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/spf13/cobra"
)

func configArgsValidator(cmd *cobra.Command, args []string) error {
	// if !len(args) % 2
	return nil
}

func configCommand(cmd *cobra.Command, args []string) {
	logger.Info("Config args length")
	logger.Info(len(args) % 2)
}
