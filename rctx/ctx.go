package rctx

import (
	"context"
)

func WithPointer[K any, V any](ctx context.Context, k K, v V) context.Context {
	x, ok := ctx.Value(k).(*V)
	if !ok {
		return context.WithValue(ctx, k, &v)
	}
	*x = v
	return ctx
}

func PointerOK[V any, K any](ctx context.Context, k K) (v *V, ok bool) {
	v, ok = ctx.Value(k).(*V)
	if !ok {
		return
	}
	return
}

func Pointer[V any, K any](ctx context.Context, k K) *V {
	v, ok := PointerOK[V, K](ctx, k)
	if !ok {
		return nil
	}
	return v
}

func WithMapValue[M any, K comparable, V any](ctx context.Context, mk M, k K, v V) context.Context {
	m, ok := ctx.Value(mk).(map[K]interface{})
	if !ok {
		m = make(map[K]interface{}, 1)
		m[k] = v
		return context.WithValue(ctx, mk, m)
	}
	m[k] = v
	return ctx
}

func MapValueOK[V any, M any, K comparable](ctx context.Context, mk M, k K) (*V, bool) {
	m, ok := ctx.Value(mk).(map[K]interface{})
	if !ok {
		return nil, false
	}
	v, ok := m[k]
	if !ok {
		return nil, false
	}
	tv, ok := v.(V)
	if !ok {
		return nil, false
	}
	return &tv, true
}

func MapValue[V any, M any, K comparable](ctx context.Context, mk M, k K) *V {
	v, ok := MapValueOK[V](ctx, mk, k)
	if !ok {
		return nil
	}
	return v
}

func Map[K comparable, M any](ctx context.Context, mk M) map[K]interface{} {
	m, ok := ctx.Value(mk).(map[K]interface{})
	if !ok {
		return nil
	}
	return m
}

func VMap[K comparable, V any, M any](ctx context.Context, mk M) map[K]V {
	m, ok := ctx.Value(mk).(map[K]interface{})
	if !ok {
		return nil
	}
	nm := make(map[K]V, len(m))
	for k, v := range m {
		tv, ok := v.(V)
		if !ok {
			continue
		}
		nm[k] = tv
	}
	return nm
}