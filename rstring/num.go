package rstring

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

func IntOrDefault(i string, d int64) int64 {
	res, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		f, err := strconv.ParseFloat(i, 64)
		if err == nil {
			return int64(f)
		}
		return d
	}
	return res
}

type Number interface {
	constraints.Integer | constraints.Float
}

func NumOrDefault[T Number](i string, d T) T {
	res, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		f, err := strconv.ParseFloat(i, 64)
		if err == nil {
			return T(f)
		}
		return d
	}
	return T(res)
}
