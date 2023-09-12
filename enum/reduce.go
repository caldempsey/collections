package enum

func Reduce[T any, R any](s []T, init R, f func(R, T) R) R {
	result := init
	for _, v := range s {
		result = f(result, v)
	}
	return result
}
