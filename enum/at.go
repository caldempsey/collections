package enum

func At[T any](s []T, index int, defaultValue T) T {
	if index < 0 || index >= len(s) {
		return defaultValue
	}
	return s[index]
}
