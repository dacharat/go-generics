package godash

func Chunk[T any](a []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	result := make([][]T, 0, len(a)/2+1)
	length := len(a)

	for i := 0; i < length; i++ {
		chunk := i / size

		if i%size == 0 {
			result = append(result, make([]T, 0, size))
		}

		result[chunk] = append(result[chunk], a[i])
	}

	return result
}
