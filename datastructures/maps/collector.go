package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
)

type Collector[K, V any] struct {
	Map Map[K, V]
}

func (collector Collector[K, V]) Supply(value misc.Pair[K, V]) {
	collector.Map.Set(value.First, value.Second)
}

func (collector Collector[K, V]) Finish() Map[K, V] {
	return collector.Map
}
