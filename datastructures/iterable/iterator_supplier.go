package iterable

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type IteratorSupplier[T any] struct {
	Iterator[T]
}

func (supplier IteratorSupplier[T]) Supply() optional.Optional[T] {
	if supplier.Valid() {
		defer supplier.Move()
		return optional.FromValue(supplier.Get())
	} else {
		return optional.Empty[T]()
	}
}
