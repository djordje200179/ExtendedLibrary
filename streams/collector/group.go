package collector

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type group[T any, K comparable] struct {
	m      map[K][]T
	mapper functions.Mapper[T, K]
}

func Group[T any, K comparable](mapper functions.Mapper[T, K]) Collector[T, map[K][]T] {
	return group[T, K]{
		m:      make(map[K][]T),
		mapper: mapper,
	}
}

func (group group[T, K]) Supply(value T) {
	key := group.mapper(value)
	group.m[key] = append(group.m[key], value)
}

func (group group[T, K]) Finish() map[K][]T {
	return group.m
}
