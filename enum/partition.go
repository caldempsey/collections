package enum

func Partition[T any](s []T, f func(T) bool) (truthy []T, falsy []T) {
	for _, v := range s {
		if f(v) {
			truthy = append(truthy, v)
		} else {
			falsy = append(falsy, v)
		}
	}
	return
}
