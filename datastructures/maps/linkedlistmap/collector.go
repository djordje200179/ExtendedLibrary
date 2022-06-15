package linkedlistmap

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[K comparable, V any] struct {
	m Map[K, V]
}

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], Map[K, V]] {
	return collector[K, V]{
		m: New[K, V](),
	}
}

func (c collector[K, V]) Supply(value misc.Pair[K, V]) {
	c.m.Set(value.First, value.Second)
}

func (c collector[K, V]) Finish() Map[K, V] {
	return c.m
}
