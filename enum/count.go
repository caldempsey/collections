package enum

func Count[T any](s []T, f func(T) bool) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}
	return count
}
