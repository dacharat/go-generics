package godash

func Includes[T comparable](a []T, t T) bool {
	for _, v := range a {
		if v == t {
			return true
		}
	}

	return false
}
