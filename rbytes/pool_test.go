package rbytes_test

import (
	"testing"

	"github.com/rikaaa0928/rtools/rbytes"
)

func TestPool(t *testing.T) {
	for i := 0; i < 1; i++ {
		bts := rbytes.Get(5)
		t.Logf("%p", bts)
		rbytes.Put(bts)
		bts = rbytes.Get(5)
		t.Logf("%p", bts)
		rbytes.Put(bts)
	}
}
