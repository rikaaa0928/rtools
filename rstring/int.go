package rstring

import "strconv"

func IntOrDefault(i string, d int64) int64 {
	res, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		return d
	}
	return res
}
