package enum

func Split[T any](s []T, count int) ([]T, []T) {
	if count < 0 {
		return []T{}, s
	}
	if count > len(s) {
		return s, []T{}
	}
	return s[:count], s[count:]
}
