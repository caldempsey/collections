package enum

func Fetch[T any](s []T, index int) T {
	if index < 0 || index >= len(s) {
		panic("index out of range")
	}
	return s[index]
}
