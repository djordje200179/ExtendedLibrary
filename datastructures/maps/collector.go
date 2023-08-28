package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
)

// Collector is a generic type that can be used to collect values into a map.
// It is used by the stream to collect values into a map.
type Collector[K, V any, M Map[K, V]] struct {
	Map M // The map to collect values into.
}

// Supply adds key-value pair to the map.
func (collector Collector[K, V, M]) Supply(value misc.Pair[K, V]) {
	collector.Map.Set(value.First, value.Second)
}

// Finish returns the map.
func (collector Collector[K, V, M]) Finish() M {
	return collector.Map
}
