package godash

func Uniq[T comparable](a []T) []T {
	result := make([]T, 0, len(a))
	seen := make(map[T]struct{}, len(a))

	for _, v := range a {
		if _, ok := seen[v]; ok {
			continue
		}

		seen[v] = struct{}{}
		result = append(result, v)
	}

	return result
}
