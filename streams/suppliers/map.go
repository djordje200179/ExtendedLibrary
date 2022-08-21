package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type mapSupplier[K comparable, V any] struct {
	m     map[K]V
	keys  []K
	index int
}

func Map[K comparable, V any](m map[K]V) Supplier[misc.Pair[K, V]] {
	keys := make([]K, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	return &mapSupplier[K, V]{m, keys, 0}
}

func (supplier *mapSupplier[K, V]) Supply() optional.Optional[misc.Pair[K, V]] {
	if supplier.index < len(supplier.keys) {
		key := supplier.keys[supplier.index]
		value := supplier.m[key]
		supplier.index++
		return optional.FromValue(misc.Pair[K, V]{key, value})
	} else {
		return optional.Empty[misc.Pair[K, V]]()
	}
}
