package godash

func Every[T any](a []T, f func(T) bool) bool {
	for _, e := range a {
		if !f(e) {
			return false
		}
	}
	return true
}
