package collections

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type supplier[T any] struct {
	Iterator[T]
}

func Supplier[T any](iterable Iterable[T]) suppliers.Supplier[T] {
	return supplier[T]{iterable.Iterator()}
}

func (supplier supplier[T]) Supply() optional.Optional[T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.Get())
	} else {
		return optional.Empty[T]()
	}
}
