package godash

// Find return first match element
func Find[T any](a []T, f func(elem T) bool) (T, bool) {
	for _, v := range a {
		if f(v) {
			return v, true
		}
	}
	var v T
	return v, false
}

func FindIndex[T any](a []T, f func(elem T) bool) int {
	for i, v := range a {
		if f(v) {
			return i
		}
	}
	return -1
}
