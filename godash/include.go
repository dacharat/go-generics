package godash

func Includes[T comparable](a []T, t T) bool {
	for _, v := range a {
		if v == t {
			return true
		}
	}

	return false
}

func IncludesBy[T comparable](a []T, f func(T) bool) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}

	return false
}
