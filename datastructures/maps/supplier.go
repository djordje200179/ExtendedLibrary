package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type supplier[K comparable, V any] struct {
	iterableSupplier suppliers.Supplier[Entry[K, V]]
}

func Supplier[K comparable, V any](m Map[K, V]) suppliers.Supplier[misc.Pair[K, V]] {
	iterableSupplier := collections.Supplier[Entry[K, V]](m)
	return supplier[K, V]{iterableSupplier}
}

func (supplier supplier[K, V]) Supply() optional.Optional[misc.Pair[K, V]] {
	if elem := supplier.iterableSupplier.Supply(); elem.HasValue() {
		entry := elem.Get()
		key, value := entry.Key(), entry.Value()
		pair := misc.Pair[K, V]{key, value}
		return optional.FromValue(pair)
	} else {
		return optional.Empty[misc.Pair[K, V]]()
	}
}
