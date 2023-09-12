package enum

type Zipped struct {
	T any
	R any
}

func Zip[T any, R any](s1 []T, s2 []R) []Zipped {
	size := len(s1)
	if len(s2) < size {
		size = len(s2)
	}
	result := make([]Zipped, size)
	for i := 0; i < size; i++ {
		result[i] = Zipped{
			T: s1[i],
			R: s2[i],
		}
	}
	return result
}
