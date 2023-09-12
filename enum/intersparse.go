package enum

func Intersperse[T any](s []T, separator T) []T {
	if len(s) == 0 {
		return []T{}
	}
	result := make([]T, 2*len(s)-1)
	for i, v := range s {
		result[2*i] = v
		if i != len(s)-1 {
			result[2*i+1] = separator
		}
	}
	return result
}
