package rctx

import (
	"context"
	"sync"
)

type waitKeyType string

var waitAllKey waitKeyType = "k"

func InitWorkerWaiter(ctx context.Context) context.Context {
	wg := &sync.WaitGroup{}
	return context.WithValue(ctx, waitAllKey, wg)
}

func WorkerAdd(ctx context.Context) {
	v := ctx.Value(waitAllKey)
	if v == nil {
		panic("init worker before add")
	}
	wg := v.(*sync.WaitGroup)
	wg.Add(1)
}

func WorkerDone(ctx context.Context) {
	v := ctx.Value(waitAllKey)
	if v == nil {
		panic("init worker before add")
	}
	wg := v.(*sync.WaitGroup)
	wg.Done()
}

func WaitAllWorker(ctx context.Context) {
	v := ctx.Value(waitAllKey)
	if v == nil {
		panic("init worker before wait")
	}
	wg := v.(*sync.WaitGroup)
	wg.Wait()
}
