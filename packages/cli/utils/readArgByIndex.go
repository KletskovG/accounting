package utils

func ReadArgByIndex(args []string, index int) string {
	if len(args) < index {
		return ""
	}

	return args[index]
}
