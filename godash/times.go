package godash

func Times[T any](n int, f func(int) T) []T {
	if n == 0 {
		return nil
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = f(i)
	}

	return result
}
