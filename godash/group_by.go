package godash

func GroupBy[T any, E comparable](a []T, f func(T) E) map[E][]T {
	result := make(map[E][]T, len(a))
	for _, v := range a {
		result[f(v)] = append(result[f(v)], v)
	}

	return result
}
