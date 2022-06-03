package godash

func Map[T, M any](a []T, f func(T, int) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e, i)
	}
	return n
}
