package enum

func Slice[T any](s []T, start int, count int) []T {
	if start < 0 || start+count > len(s) {
		return []T{}
	}
	return s[start : start+count]
}
