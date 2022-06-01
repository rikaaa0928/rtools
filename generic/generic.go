package generic

func TakeAddr[T any](v T) *T {
	return &v
}
