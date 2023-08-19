package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
)

type Collector[K, V any, M Map[K, V]] struct {
	Map M
}

func (collector Collector[K, V, M]) Supply(value misc.Pair[K, V]) {
	collector.Map.Set(value.First, value.Second)
}

func (collector Collector[K, V, M]) Finish() M {
	return collector.Map
}
