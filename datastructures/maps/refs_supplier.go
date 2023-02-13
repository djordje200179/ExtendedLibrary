package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type RefsSupplier[K comparable, V any] struct {
	Iterator[K, V]
}

func (supplier RefsSupplier[K, V]) Supply() optional.Optional[misc.Pair[K, *V]] {
	if !supplier.Iterator.Valid() {
		return optional.Empty[misc.Pair[K, *V]]()
	}

	defer supplier.Iterator.Move()

	key := supplier.Iterator.Key()
	valueRef := supplier.Iterator.ValueRef()
	pair := misc.Pair[K, *V]{key, valueRef}

	return optional.FromValue(pair)
}
