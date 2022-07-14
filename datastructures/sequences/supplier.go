package sequences

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type refsSupplier[T any] struct {
	Iterator[T]
}

func RefsSupplier[T any](sequence Sequence[T]) suppliers.Supplier[*T] {
	return refsSupplier[T]{sequence.ModifyingIterator()}
}

func (supplier refsSupplier[T]) Supply() optional.Optional[*T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.GetRef())
	} else {
		return optional.Empty[*T]()
	}
}
