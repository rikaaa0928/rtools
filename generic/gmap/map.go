package gmap

func GetWithDefault[K comparable, V any](m map[K]V, key K, dft V) V {
	t, ok := m[key]
	if ok {
		return t
	}
	return dft
}
