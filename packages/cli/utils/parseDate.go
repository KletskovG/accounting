package utils

import (
	"time"

	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
)

func ParseDate(input string) string {
	result, err := time.Parse(common.DateLayout, input)

	if err != nil {
		logger.Info("Cant parse date, ", input, err)
		return ""
	}

	return result.Format(common.DateLayout)
}
