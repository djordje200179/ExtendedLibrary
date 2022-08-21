package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type Supplier[K comparable, V any] struct {
	collections.Iterator[Entry[K, V]]
}

func ValuesSupplier[K comparable, V any](m Map[K, V]) suppliers.Supplier[misc.Pair[K, V]] {
	supplier := Supplier[K, V]{m.Iterator()}
	return suppliers.FunctionSupplier[misc.Pair[K, V]](supplier.NextValue)
}

func RefsSupplier[K comparable, V any](m Map[K, V]) suppliers.Supplier[misc.Pair[K, *V]] {
	supplier := Supplier[K, V]{m.Iterator()}
	return suppliers.FunctionSupplier[misc.Pair[K, *V]](supplier.NextRef)
}

func (supplier Supplier[K, V]) NextValue() optional.Optional[misc.Pair[K, V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, V]]()
	}

	defer supplier.Iterator.Move()

	entry := supplier.Iterator.Get()
	data := misc.Pair[K, V]{entry.Key(), entry.Value()}

	return optional.FromValue(data)
}

func (supplier Supplier[K, V]) NextRef() optional.Optional[misc.Pair[K, *V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, *V]]()
	}

	defer supplier.Iterator.Move()

	entry := supplier.Iterator.Get()
	data := misc.Pair[K, *V]{entry.Key(), entry.ValueRef()}

	return optional.FromValue(data)
}
