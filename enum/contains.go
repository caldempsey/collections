package enum

// Contains function
func Contains[T comparable](s []T, el T) bool {
	for _, v := range s {
		if v == el {
			return true
		}
	}
	return false
}
