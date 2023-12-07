package utils

import (
	"time"

	"github.com/kletskovg/accounting/packages/logger"
)

const dateLayout = "2006-01-02"

func ParseDate(input string) string {
	result, err := time.Parse(dateLayout, input)

	if err != nil {
		logger.Error("Cant parse date, ", input, err)
	}

	return result.Format(dateLayout)
}
