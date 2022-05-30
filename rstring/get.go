package rstring

func GetFirstAvaliable(values ...string) string {
	for _, v := range values {
		if len(v) > 0 {
			return v
		}
	}
	return ""
}
