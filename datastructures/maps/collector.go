package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[K comparable, V any] struct {
	m Map[K, V]
}

func Collector[K comparable, V any](empty Map[K, V]) streams.Collector[misc.Pair[K, V], Map[K, V]] {
	return collector[K, V]{empty}
}

func (collector collector[K, V]) Supply(value misc.Pair[K, V]) {
	collector.m.Set(value.First, value.Second)
}

func (collector collector[K, V]) Finish() Map[K, V] { return collector.m }
