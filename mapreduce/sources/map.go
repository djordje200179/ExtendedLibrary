package sources

import (
	"github.com/djordje200179/extendedlibrary/mapreduce"
	"github.com/djordje200179/extendedlibrary/misc"
)

func NewMapSource[K comparable, V any](m map[K]V) mapreduce.Source[K, V] {
	source := make(chan misc.Pair[K, V], 100)

	go func() {
		for key, value := range m {
			source <- misc.Pair[K, V]{key, value}
		}
		close(source)
	}()

	return source
}
