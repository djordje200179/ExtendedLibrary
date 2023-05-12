package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

type sliceIterator[T any] struct {
	slice []T
	index int
}

func SliceValues[T any](slice []T) streams.Supplier[T] {
	iterator := sliceIterator[T]{slice, 0}
	return iterator.nextValue
}

func SliceRefs[T any](slice []T) streams.Supplier[*T] {
	iterator := sliceIterator[T]{slice, 0}
	return iterator.nextRef
}

func (supplier *sliceIterator[T]) nextValue() optional.Optional[T] {
	if supplier.index >= len(supplier.slice) {
		return optional.Empty[T]()
	}

	curr := supplier.slice[supplier.index]
	supplier.index++
	return optional.FromValue(curr)
}

func (supplier *sliceIterator[T]) nextRef() optional.Optional[*T] {
	if supplier.index >= len(supplier.slice) {
		return optional.Empty[*T]()
	}

	curr := &supplier.slice[supplier.index]
	supplier.index++
	return optional.FromValue(curr)
}
