package enum

func Max[T any](s []T, compare func(a, b T) bool) T {
	if len(s) == 0 {
		panic("slice is empty")
	}
	max := s[0]
	for _, v := range s[1:] {
		if compare(v, max) {
			max = v
		}
	}
	return max
}
