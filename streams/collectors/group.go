package collectors

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type groupCollector[T any, K comparable] struct {
	m      map[K][]T
	mapper functions.Mapper[T, K]
}

func Group[T any, K comparable](mapper functions.Mapper[T, K]) streams.Collector[T, map[K][]T] {
	return groupCollector[T, K]{
		m:      make(map[K][]T),
		mapper: mapper,
	}
}

func (collector groupCollector[T, K]) Supply(value T) {
	key := collector.mapper(value)
	collector.m[key] = append(collector.m[key], value)
}

func (collector groupCollector[T, K]) Finish() map[K][]T { return collector.m }
