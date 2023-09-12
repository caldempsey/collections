package enum

// Dedup assumes that the input slice is sorted.
func Dedup[T comparable](s []T) []T {
	if len(s) == 0 {
		return []T{}
	}
	result := []T{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			result = append(result, s[i])
		}
	}
	return result
}
