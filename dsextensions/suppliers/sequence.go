package suppliers

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type sequenceRefsSupplier[T any] struct {
	sequences.Iterator[T]
}

func FromSequenceRefs[T any](sequence sequences.Sequence[T]) suppliers.Supplier[*T] {
	return sequenceRefsSupplier[T]{sequence.ModifyingIterator()}
}

func (supplier sequenceRefsSupplier[T]) Supply() optional.Optional[*T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.GetRef())
	} else {
		return optional.Empty[*T]()
	}
}
