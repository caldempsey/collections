package enum

import "sort"

func Sort[T any](s []T, less func(a, b T) bool) {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
}
