package gslice

func GetMap[K comparable, V any](slice []V, keyFunc func(V) K) map[K]V {
	res := make(map[K]V, len(slice))
	for _, v := range slice {
		k := keyFunc(v)
		res[k] = v
	}
	return res
}
