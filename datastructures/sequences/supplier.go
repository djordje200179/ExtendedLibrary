package sequences

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type Supplier[T any] struct {
	Iterator[T]
}

func ValuesSupplier[T any](sequence Sequence[T]) suppliers.Supplier[T] {
	supplier := Supplier[T]{sequence.ModifyingIterator()}
	return suppliers.FunctionSupplier[T](supplier.NextValue)
}

func RefsSupplier[T any](sequence Sequence[T]) suppliers.Supplier[*T] {
	supplier := Supplier[T]{sequence.ModifyingIterator()}
	return suppliers.FunctionSupplier[*T](supplier.NextRef)
}

func (supplier Supplier[T]) NextValue() optional.Optional[T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.Get())
	} else {
		return optional.Empty[T]()
	}
}

func (supplier Supplier[T]) NextRef() optional.Optional[*T] {
	it := supplier.Iterator
	if it.Valid() {
		defer it.Move()
		return optional.FromValue(it.GetRef())
	} else {
		return optional.Empty[*T]()
	}
}
