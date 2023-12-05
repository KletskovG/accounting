package utils

import "github.com/kletskovg/accounting/packages/logger"

func ReadArgByIndex(args []string, index int) string {
	if len(args) <= index {
		logger.Info("trying to use index ", index, "from ", args)
		return ""
	}

	return args[index]
}
