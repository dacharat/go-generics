package godash

func Partition[T any](collection []T, f func(T) bool) ([]T, []T) {
	trust := make([]T, 0, len(collection))
	untrust := make([]T, 0, len(collection))

	for _, c := range collection {
		if f(c) {
			trust = append(trust, c)
		} else {
			untrust = append(untrust, c)
		}
	}

	return trust, untrust
}
