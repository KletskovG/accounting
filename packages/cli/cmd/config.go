package cmd

import (
	"os"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
	"github.com/spf13/cobra"
)

func configCommand(cmd *cobra.Command, args []string) {
	if err := os.Remove(common.CliUserConfigPath); err != nil {
		logger.Info("Cant remove config file %s", err)
	}
	logger.Info("Config file was cleared, run any command to reconfigure")
}
