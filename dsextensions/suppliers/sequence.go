package suppliers

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type sequenceSupplier[T any] struct {
	sequences.Iterator[T]
}

func FromSequence[T any](sequence sequences.Sequence[T]) suppliers.Supplier[*T] {
	return sequenceSupplier[T]{sequence.ModifyingIterator()}
}

func (supplier sequenceSupplier[T]) Supply() optional.Optional[*T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.GetRef())
	} else {
		return optional.Empty[*T]()
	}
}
