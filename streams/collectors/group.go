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

func (c groupCollector[T, K]) Supply(value T) {
	key := c.mapper(value)
	c.m[key] = append(c.m[key], value)
}

func (c groupCollector[T, K]) Finish() map[K][]T {
	return c.m
}

func Group2[K comparable, V any](s streams.Stream2[K, V]) map[K][]V {
	m := make(map[K][]V)

	for k, v := range s {
		m[k] = append(m[k], v)
	}

	return m
}
