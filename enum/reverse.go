package enum

func Reverse[T any](s []T) []T {
	size := len(s)
	result := make([]T, size)
	for i, v := range s {
		result[size-1-i] = v
	}
	return result
}
