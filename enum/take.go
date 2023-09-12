package enum

func Take[T any](s []T, count int) []T {
	if count < 0 {
		return []T{}
	}
	if count > len(s) {
		return s
	}
	return s[:count]
}
