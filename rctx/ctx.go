package rctx

import (
	"context"
)

var defaultMKey interface{}

type mapWithIKey[K comparable, V any] struct {
	iKey interface{}
}

func (mk *mapWithIKey[K, V]) WithMapValue(ctx context.Context, k K, v V) context.Context {
	m, ok := ctx.Value(mk.iKey).(map[K]interface{})
	if !ok {
		m = make(map[K]interface{}, 1)
		m[k] = v
		return context.WithValue(ctx, mk.iKey, m)
	}
	m[k] = v
	return ctx
}

func (mk *mapWithIKey[K, V]) MapValueOK(ctx context.Context, k K) (*V, bool) {
	val := ctx.Value(mk.iKey)
	m, ok := val.(map[K]interface{})
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

func (mk *mapWithIKey[K, V]) MapValue(ctx context.Context, k K) *V {
	v, ok := mk.MapValueOK(ctx, k)
	if !ok {
		return nil
	}
	return v
}

func (mk *mapWithIKey[K, V]) Map(ctx context.Context) map[K]interface{} {
	m, ok := ctx.Value(mk.iKey).(map[K]interface{})
	if !ok {
		return nil
	}
	return m
}

func (mk *mapWithIKey[K, V]) TypedMap(ctx context.Context) map[K]V {
	m, ok := ctx.Value(mk.iKey).(map[K]interface{})
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

type MapWithIndexKey[K comparable, V any] interface {
	WithMapValue(ctx context.Context, k K, v V) context.Context
	MapValueOK(ctx context.Context, k K) (*V, bool)
	MapValue(ctx context.Context, k K) *V
	Map(ctx context.Context) map[K]interface{}
	TypedMap(ctx context.Context) map[K]V
}

func MapKey[K comparable, V any](mk interface{}) MapWithIndexKey[K, V] {
	return &mapWithIKey[K, V]{mk}
}

func DefaultMapKey(k interface{}) {
	defaultMKey = k
}

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

//func DWithMapValue[K comparable, V any](ctx context.Context, k K, v V) context.Context {
//	return WithMapValue[any, K, V](ctx, defaultMKey, k, v)
//}

func WithMapValue[K comparable, V any](ctx context.Context, k K, v V) context.Context {
	return MapKey[K, V](defaultMKey).WithMapValue(ctx, k, v)
}

//func DMapValueOK[V any, K comparable](ctx context.Context, k K) (*V, bool) {
//	return MapValueOK[V, any, K](ctx, defaultMKey, k)
//}

func MapValueOK[V any, K comparable](ctx context.Context, k K) (*V, bool) {
	return MapKey[K, V](defaultMKey).MapValueOK(ctx, k)
}

//func DMapValue[V any, K comparable](ctx context.Context, k K) *V {
//	return MapValue[V, any, K](ctx, defaultMKey, k)
//}

func MapValue[V any, K comparable](ctx context.Context, k K) *V {
	return MapKey[K, V](defaultMKey).MapValue(ctx, k)
}

//func DMap[K comparable](ctx context.Context) map[K]interface{} {
//	return Map[K, any](ctx, defaultMKey)
//}

func Map[K comparable](ctx context.Context) map[K]interface{} {
	return MapKey[K, interface{}](defaultMKey).Map(ctx)
}

//func DVMap[K comparable, V any](ctx context.Context) map[K]V {
//	return VMap[K, V, any](ctx, defaultMKey)
//}

func TypedMap[K comparable, V any](ctx context.Context) map[K]V {
	return MapKey[K, V](defaultMKey).TypedMap(ctx)
}
