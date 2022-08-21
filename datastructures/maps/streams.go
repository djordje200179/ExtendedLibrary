package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

type supplier[K comparable, V any] struct {
	collections.Iterator[Entry[K, V]]
}

func ValuesStream[K comparable, V any](m Map[K, V]) streams.Stream[misc.Pair[K, V]] {
	supplier := supplier[K, V]{m.Iterator()}
	return streams.FromFiniteGenerator(supplier.NextValue)
}

func RefsStream[K comparable, V any](m Map[K, V]) streams.Stream[misc.Pair[K, *V]] {
	supplier := supplier[K, V]{m.Iterator()}
	return streams.FromFiniteGenerator(supplier.NextRef)
}

func (supplier supplier[K, V]) NextValue() optional.Optional[misc.Pair[K, V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, V]]()
	}

	defer supplier.Iterator.Move()

	entry := supplier.Iterator.Get()
	data := misc.Pair[K, V]{entry.Key(), entry.Value()}

	return optional.FromValue(data)
}

func (supplier supplier[K, V]) NextRef() optional.Optional[misc.Pair[K, *V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, *V]]()
	}

	defer supplier.Iterator.Move()

	entry := supplier.Iterator.Get()
	data := misc.Pair[K, *V]{entry.Key(), entry.ValueRef()}

	return optional.FromValue(data)
}