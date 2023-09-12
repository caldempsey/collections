package enum

func Drop[T any](s []T, count int) []T {
	if count < 0 {
		return s
	}
	if count > len(s) {
		return []T{}
	}
	return s[count:]
}
