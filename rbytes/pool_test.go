package rbytes_test

import (
	"testing"

	"github.com/rikaaa0928/rtools/rbytes"
)

func TestPool(t *testing.T) {
	for i := 0; i < 5; i++ {
		bts := rbytes.Get(5)
		t.Logf("%p", bts)
		t.Logf("%d", bts[0])
		bts[0] = byte(i)
		rbytes.Put(bts)
		bts = rbytes.Get(3)
		t.Logf("%p", bts)
		t.Logf("%d", bts[0])
		bts[0] = byte(i + 10)
		rbytes.Put(bts)
	}
}
