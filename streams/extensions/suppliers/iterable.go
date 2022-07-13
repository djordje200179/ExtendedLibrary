package suppliers

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type iterableSupplier[T any] struct {
	datastructures.Iterator[T]
}

func FromIterable[T any](iterable datastructures.Iterable[T]) suppliers.Supplier[T] {
	return iterableSupplier[T]{iterable.Iterator()}
}

func (supplier iterableSupplier[T]) Supply() optional.Optional[T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.Get())
	} else {
		return optional.Empty[T]()
	}
}
