package collectors

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps/linkedlistmap"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type group[T any, K comparable] struct {
	m      linkedlistmap.Map[K, *array.Array[T]]
	mapper functions.Mapper[T, K]
}

func Group[T any, K comparable](mapper functions.Mapper[T, K]) streams.Collector[T, linkedlistmap.Map[K, *array.Array[T]]] {
	return group[T, K]{
		m:      linkedlistmap.New[K, *array.Array[T]](),
		mapper: mapper,
	}
}

func (group group[T, K]) Supply(value T) {
	key := group.mapper(value)

	if !group.m.Contains(key) {
		group.m.Set(key, array.New[T](0))
	}

	arr := group.m.Get(key)
	arr.Append(value)
}

func (group group[T, K]) Finish() linkedlistmap.Map[K, *array.Array[T]] {
	return group.m
}

func Partition[T any](predictor functions.Predictor[T]) streams.Collector[T, linkedlistmap.Map[bool, *array.Array[T]]] {
	return Group[T](functions.Mapper[T, bool](predictor))
}
