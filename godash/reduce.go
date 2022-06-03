package godash

func Reduce[T, M any](a []T, f func(M, T) M, init M) M {
	i := init
	for _, e := range a {
		i = f(i, e)
	}
	return i
}
