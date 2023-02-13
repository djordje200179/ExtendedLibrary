package collections

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type RefsSupplier[T any] struct {
	Iterator[T]
}

func (supplier RefsSupplier[T]) Supply() optional.Optional[*T] {
	if supplier.Valid() {
		defer supplier.Move()
		return optional.FromValue(supplier.GetRef())
	} else {
		return optional.Empty[*T]()
	}
}
