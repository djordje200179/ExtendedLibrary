package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

type supplier[K comparable, V any] struct {
	Iterator[K, V]
}

func ValuesStream[K comparable, V any](m Map[K, V]) streams.Stream[misc.Pair[K, V]] {
	supplier := supplier[K, V]{m.ModifyingIterator()}
	return streams.FromFiniteGenerator(supplier.NextValue)
}

func RefsStream[K comparable, V any](m Map[K, V]) streams.Stream[misc.Pair[K, *V]] {
	supplier := supplier[K, V]{m.ModifyingIterator()}
	return streams.FromFiniteGenerator(supplier.NextRef)
}

func (supplier supplier[K, V]) NextValue() optional.Optional[misc.Pair[K, V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, V]]()
	}

	defer supplier.Iterator.Move()

	return optional.FromValue(supplier.Iterator.Get())
}

func (supplier supplier[K, V]) NextRef() optional.Optional[misc.Pair[K, *V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, *V]]()
	}

	defer supplier.Iterator.Move()

	key := supplier.Iterator.Key()
	value := supplier.Iterator.ValueRef()
	pair := misc.Pair[K, *V]{key, value}

	return optional.FromValue(pair)
}
