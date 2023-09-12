package enum

func GroupBy[T any, K comparable](s []T, f func(T) K) map[K][]T {
	m := make(map[K][]T)
	for _, v := range s {
		key := f(v)
		m[key] = append(m[key], v)
	}
	return m
}
