package common

func MapSlice[T any, R any](slice []T, mapFn func(el T) R) []R {
	copy := make([]R, 0)

	for _, item := range slice {
		copy = append(copy, mapFn(item))
	}

	return copy
}
