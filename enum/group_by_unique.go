package enum

// GroupByUnique will group to a key only once, e.g. for unique keys
func GroupByUnique[T any, K comparable](s []T, f func(T) K) map[K]T {
	m := make(map[K]T)
	for _, v := range s {
		key := f(v)
		m[key] = v
	}
	return m
}
