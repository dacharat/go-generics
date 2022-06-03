package godash

func ForEach[T any](a []T, f func(T, int)) {
	for i, v := range a {
		f(v, i)
	}
}
