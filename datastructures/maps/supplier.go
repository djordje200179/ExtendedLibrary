package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

func ValueSupplier[K comparable, V any](m Map[K, V]) functions.EmptyGenerator[optional.Optional[misc.Pair[K, V]]] {
	iterator := m.Iterator()

	return func() optional.Optional[misc.Pair[K, V]] {
		if !iterator.Valid() {
			return optional.Empty[misc.Pair[K, V]]()
		}

		defer iterator.Move()

		entry := iterator.Get()
		data := misc.Pair[K, V]{entry.Key(), entry.Value()}

		return optional.FromValue(data)
	}
}

func RefSupplier[K comparable, V any](m Map[K, V]) functions.EmptyGenerator[optional.Optional[misc.Pair[K, *V]]] {
	iterator := m.Iterator()

	return func() optional.Optional[misc.Pair[K, *V]] {
		if !iterator.Valid() {
			return optional.Empty[misc.Pair[K, *V]]()
		}

		defer iterator.Move()

		entry := iterator.Get()
		data := misc.Pair[K, *V]{entry.Key(), entry.ValueRef()}

		return optional.FromValue(data)
	}
}
