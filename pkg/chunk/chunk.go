package chunk

// SliceToChunks separate slice to chunks with chunkSize max size.
func SliceToChunks[T any](array []T, chunkSize int) [][]T {
	var result [][]T
	arrayLen := len(array)

	for i := 0; i < arrayLen; i += chunkSize {
		j := i + chunkSize
		if j > arrayLen {
			j = arrayLen
		}
		result = append(result, array[i:j])
	}
	return result
}

// ChanToChunks separate chan array to chunks with chunkSize max size.
func ChanToChunks[T any](ch <-chan T, chunkSize int) [][]T {
	var result [][]T //nolint:prealloc
	var chunk []T    //nolint:prealloc
	for item := range ch {
		chunk = append(chunk, item)
		if len(chunk) < chunkSize {
			continue
		}

		result = append(result, chunk)
		chunk = []T{}
	}

	if len(chunk) != 0 {
		result = append(result, chunk)
	}

	return result
}
