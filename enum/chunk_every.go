package enum

func ChunkEvery[T any](s []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	var result [][]T
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[i:end])
	}
	return result
}
