package enum

func Reject[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if !f(v) {
			result = append(result, v)
		}
	}
	return result
}
