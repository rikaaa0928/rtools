package gslice

func InSlice[T comparable](set []T, obj T) bool {
	for _, v := range set {
		if v == obj {
			return true
		}
	}
	return false
}
