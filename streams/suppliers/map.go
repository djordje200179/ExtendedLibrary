package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

type mapIterator[K comparable, V any] struct {
	m     map[K]V
	keys  []K
	index int
}

func Map[K comparable, V any](m map[K]V) streams.Supplier[misc.Pair[K, V]] {
	keys := make([]K, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	iterator := &mapIterator[K, V]{m, keys, 0}

	return iterator.supply
}

func (iterator *mapIterator[K, V]) supply() optional.Optional[misc.Pair[K, V]] {
	if iterator.index < len(iterator.keys) {
		key := iterator.keys[iterator.index]
		value := iterator.m[key]
		iterator.index++
		return optional.FromValue(misc.Pair[K, V]{key, value})
	} else {
		return optional.Empty[misc.Pair[K, V]]()
	}
}
