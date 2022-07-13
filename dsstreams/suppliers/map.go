package suppliers

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type mapSupplier[K comparable, V any] iterableSupplier[maps.Entry[K, V]]

func FromMap[K comparable, V any](m maps.Map[K, V]) suppliers.Supplier[misc.Pair[K, V]] {
	return mapSupplier[K, V](FromIterable[maps.Entry[K, V]](m).(iterableSupplier[maps.Entry[K, V]]))
}

func (supplier mapSupplier[K, V]) Supply() optional.Optional[misc.Pair[K, V]] {
	if elem := iterableSupplier[maps.Entry[K, V]](supplier).Supply(); elem.HasValue() {
		entry := elem.Get()
		key, value := entry.Key(), entry.Value()
		pair := misc.Pair[K, V]{key, value}
		return optional.FromValue(pair)
	} else {
		return optional.Empty[misc.Pair[K, V]]()
	}
}
