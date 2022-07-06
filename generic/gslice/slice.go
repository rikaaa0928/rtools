package gslice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func InSlice[T comparable](slice []T, obj T) bool {
	for _, v := range slice {
		if v == obj {
			return true
		}
	}
	return false
}

func Sort[T constraints.Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func CompareOrdered[T constraints.Ordered](slice, slice2 []T) bool {
	Sort(slice)
	Sort(slice2)
	return Compare(slice, slice2)
}

func Compare[T comparable](slice, slice2 []T) bool {
	if len(slice) != len(slice2) {
		return false
	}
	for i := range slice {
		if slice[i] != slice2[i] {
			return false
		}
	}
	return true
}
