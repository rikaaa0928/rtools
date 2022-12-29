package rstring_test

import (
	"github.com/rikaaa0928/rtools/rstring"
	"testing"
)

func TestNumOrDefault(t *testing.T) {
	var x int32 = rstring.NumOrDefault("1", int32(2))
	t.Log(x)
	var y int = rstring.NumOrDefault("1.1", 2)
	t.Log(y)
	var f float64 = rstring.NumOrDefault("1", 3.0)
	t.Log(f)
	var f2 float32 = rstring.NumOrDefault("3.3", float32(4.3))
	t.Log(f2)
}
