package enum

func Concat[T any](s1 []T, s2 []T) []T {
	return append(s1, s2...)
}
