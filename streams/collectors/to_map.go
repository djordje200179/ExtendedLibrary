package collectors

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type mapCollector[K comparable, V any] map[K]V

func ToMap[K comparable, V any]() streams.Collector[misc.Pair[K, V], map[K]V] {
	return mapCollector[K, V](make(map[K]V))
}

func (collector mapCollector[K, V]) Supply(pair misc.Pair[K, V]) {
	collector[pair.First] = pair.Second
}

func (collector mapCollector[K, V]) Finish() map[K]V {
	return collector
}
