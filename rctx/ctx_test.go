package rctx_test

import (
	"context"
	"encoding/json"
	"github.com/rikaaa0928/rtools/rctx"
	"testing"
	"time"
)

func TestPointer(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	ctx = rctx.WithPointer(ctx, "k2", "vv1")
	t.Log(ctx.Value("k1"))
	t.Log(*rctx.Pointer[string](ctx, "k2"))
	t.Log(rctx.Pointer[string](ctx, "kk2"))
	testPointer(ctx, t)
	t.Log(ctx.Value("k1"))
	t.Log(*ctx.Value("k2").(*string))
}

func testPointer(ctx context.Context, t *testing.T) {
	ctx = context.WithValue(ctx, "k1", "v2")
	t.Log(ctx.Value("k1"))
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
	ctx = rctx.MapKey[string, *structX]("k1").WithMapValue(ctx, "k2", x)
	ctx = rctx.MapKey[string, interX]("k3").WithMapValue(ctx, "k4", y)
	bytes, err := json.Marshal(rctx.MapKey[string, interface{}]("k1").Map(ctx))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	bytes, err = json.Marshal(rctx.MapKey[string, interface{}]("k3").Map(ctx))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
	xxv := rctx.MapKey[string, *structX]("k1").MapValue(ctx, "k2")
	t.Logf("%v,%v,%v", xxv, *xxv, **xxv)
	xxy := *rctx.MapKey[string, interX]("k3").MapValue(ctx, "k4")
	t.Logf("%v", xxy.A())
}

func TestConvert(t *testing.T) {
	var x interface{} = make(map[string]string)
	_, ok := x.(map[string]interface{})
	t.Log(ok)
}

func TestEmptyStruct(t *testing.T) {
	type a string
	ctx := context.Background()
	k := struct {
	}{}
	ctx = context.WithValue(ctx, k, "v1")
	t.Log(ctx.Value(k))
	t.Log(ctx.Value(struct {
	}{}))
	var kk a = "k"
	ctx = context.WithValue(ctx, kk, "v2")
	t.Log(ctx.Value(kk))
	t.Log(ctx.Value("k"))
	ctx = rctx.WithWorkerWaiter(ctx)
	for i := 0; i < 5; i++ {
		go func(i int) {
			for j := 0; j < 5; j++ {
				rctx.WorkerAdd(ctx)
				go func(i, j int) {
					time.Sleep(time.Second * time.Duration(i+j))
					t.Logf("worker %d:%d done\n", i, j)
					rctx.WorkerDone(ctx)
				}(i, j)
			}
		}(i)
	}
	time.Sleep(time.Second)
	rctx.WaitAllWorker(ctx)
	t.Log("done")
}

func TestMapValue(t *testing.T) {
	type a string
	ctx := context.Background()
	var kk a = "k"
	ctx = rctx.MapKey[string, string](kk).WithMapValue(ctx, "k1", "v2")
	t.Log(rctx.MapKey[string, string]("k").MapValue(ctx, "k1"))
	t.Log(*rctx.MapKey[string, interface{}](a("k")).MapValue(ctx, "k1"))
	t.Log(*rctx.MapKey[string, string](kk).MapValue(ctx, "k1"))
	var k2 interface{} = kk
	t.Log(*rctx.MapKey[string, string](k2).MapValue(ctx, "k1"))
	rctx.DefaultMapKey(kk)
	t.Log(*rctx.MapValue[string, string](ctx, "k1"))
	ctx = rctx.WithMapValue(ctx, "k2", "vvv")
	t.Log(rctx.TypedMap[string, string](ctx))
	t.Log(rctx.Map[string](ctx))
}
