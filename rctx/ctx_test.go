package rctx_test

import (
	"context"
	"encoding/json"
	"github.com/rikaaa0928/rtools/rctx"
	"testing"
)

func TestPointer(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	ctx = rctx.WithPointer(ctx, "k2", "vv1")
	ctx = rctx.WithMapValue(ctx, "k3", "k4", "vvv1")
	t.Log(ctx.Value("k1"))
	t.Log(*rctx.Pointer[string](ctx, "k2"))
	t.Log(rctx.Pointer[string](ctx, "kk2"))
	t.Log(*rctx.MapValue[string](ctx, "k3", "k4"))
	testPointer(ctx, t)
	t.Log(ctx.Value("k1"))
	t.Log(*ctx.Value("k2").(*string))
	t.Log(*rctx.MapValue[string](ctx, "k3", "k4"))
	t.Log(*rctx.MapValue[string](ctx, "k3", "k5"))
	t.Log(*rctx.MapValue[int](ctx, "k3", "k6"))
	bytes, err := json.Marshal(rctx.Map[string](ctx, "k3"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	bytes, err = json.Marshal(rctx.VMap[string, string](ctx, "k3"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
}

func testPointer(ctx context.Context, t *testing.T) {
	ctx = context.WithValue(ctx, "k1", "v2")
	ctx = rctx.WithPointer(ctx, "k2", "vv2")
	ctx = rctx.WithMapValue(ctx, "k3", "k4", "vvv1")
	ctx = rctx.WithMapValue(ctx, "k3", "k5", "vvvx")
	ctx = rctx.WithMapValue(ctx, "k3", "k6", 6)
	t.Log(ctx.Value("k1"))
	t.Log(*rctx.Pointer[string](ctx, "k2"))
	t.Log(*rctx.MapValue[string](ctx, "k3", "k4"))
	t.Log(*rctx.MapValue[string](ctx, "k3", "k5"))
	t.Log(*rctx.MapValue[int](ctx, "k3", "k6"))
}

type interX interface {
	A() int
}

type structX struct {
	X int
}

func (s *structX) A() int {
	return s.X
}

func TestInterface(t *testing.T) {
	ctx := context.Background()
	var x *structX = &structX{X: 1}
	var y interX = &structX{X: 2}
	ctx = rctx.WithMapValue(ctx, "k1", "k2", x)
	ctx = rctx.WithMapValue(ctx, "k3", "k4", y)
	bytes, err := json.Marshal(rctx.Map[string](ctx, "k1"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	bytes, err = json.Marshal(rctx.Map[string](ctx, "k3"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	xxv := rctx.MapValue[*structX](ctx, "k1", "k2")
	t.Logf("%v,%v,%v", xxv, *xxv, **xxv)
	xxy := *rctx.MapValue[interX](ctx, "k3", "k4")
	t.Logf("%v", xxy.A())
}

func TestConvert(t *testing.T) {
	var x interface{} = make(map[string]string)
	_, ok := x.(map[string]interface{})
	t.Log(ok)
}