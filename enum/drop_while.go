package enum

func DropWhile[T any](s []T, f func(T) bool) []T {
	for i, v := range s {
		if !f(v) {
			return s[i:]
		}
	}
	return []T{}
}
