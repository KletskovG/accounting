package common

func ReduceSlice[T any, R any](slice []R, reduceFn func(acc T, curr R) T, defaultValue T) T {
	result := defaultValue

	for _, value := range slice {
		result = reduceFn(result, value)
	}

	return result
}
