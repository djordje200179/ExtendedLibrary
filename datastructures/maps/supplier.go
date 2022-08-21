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

func makeSupplier[K comparable, V any](m Map[K, V]) supplier[K, V] {
	return supplier[K, V]{collections.Supplier[Entry[K, V]](m)}
}

func (supplier supplier[K, V]) nextEntry() optional.Optional[Entry[K, V]] {
	if elem := supplier.iterableSupplier.Supply(); elem.HasValue() {
		return optional.FromValue(elem.Get())
	} else {
		return optional.Empty[Entry[K, V]]()
	}
}

type valueSupplier[K comparable, V any] struct {
	supplier[K, V]
}

func (supplier valueSupplier[K, V]) Supply() optional.Optional[misc.Pair[K, V]] {
	entry := supplier.nextEntry()
	if entry.HasValue() {
		data := entry.Get()
		key, value := data.Key(), data.Value()
		return optional.FromValue(misc.Pair[K, V]{key, value})
	} else {
		return optional.Empty[misc.Pair[K, V]]()
	}
}

func ValueSupplier[K comparable, V any](m Map[K, V]) suppliers.Supplier[misc.Pair[K, V]] {
	return valueSupplier[K, V]{makeSupplier(m)}
}

type refSupplier[K comparable, V any] struct {
	supplier[K, V]
}

func RefSupplier[K comparable, V any](m Map[K, V]) suppliers.Supplier[misc.Pair[K, *V]] {
	return refSupplier[K, V]{makeSupplier(m)}
}

func (supplier refSupplier[K, V]) Supply() optional.Optional[misc.Pair[K, *V]] {
	entry := supplier.nextEntry()
	if entry.HasValue() {
		data := entry.Get()
		return optional.FromValue(misc.Pair[K, *V]{data.Key(), data.ValueRef()})
	} else {
		return optional.Empty[misc.Pair[K, *V]]()
	}
}
