package collectors

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps/linkedlistmap"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type groupCollector[T any, K comparable] struct {
	m      linkedlistmap.Map[K, *array.Array[T]]
	mapper functions.Mapper[T, K]
}

func Group[T any, K comparable](mapper functions.Mapper[T, K]) streams.Collector[T, linkedlistmap.Map[K, *array.Array[T]]] {
	return groupCollector[T, K]{
		m:      linkedlistmap.New[K, *array.Array[T]](),
		mapper: mapper,
	}
}

func (collector groupCollector[T, K]) Supply(value T) {
	key := collector.mapper(value)

	if !collector.m.Contains(key) {
		collector.m.Set(key, array.New[T]())
	}

	arr := collector.m.Get(key)
	arr.Append(value)
}

func (collector groupCollector[T, K]) Finish() linkedlistmap.Map[K, *array.Array[T]] {
	return collector.m
}

func Partition[T any](predictor functions.Predictor[T]) streams.Collector[T, linkedlistmap.Map[bool, *array.Array[T]]] {
	return Group[T](functions.Mapper[T, bool](predictor))
}
