package enum

func Min[T any](s []T, compare func(a, b T) bool) T {
	if len(s) == 0 {
		panic("slice is empty")
	}
	min := s[0]
	for _, v := range s[1:] {
		if compare(v, min) {
			min = v
		}
	}
	return min
}
