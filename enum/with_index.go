package enum

type AnyWithIndex struct {
	T any
	I int
}

func WithIndex[T any](s []T) []AnyWithIndex {
	result := make([]AnyWithIndex, len(s))
	for i, v := range s {
		result[i] = AnyWithIndex{v, i}
	}
	return result
}
