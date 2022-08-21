package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type sliceSupplier[T any] struct {
	slice []T
	index int
}

func (supplier *sliceSupplier[T]) NextValue() optional.Optional[T] {
	if supplier.index >= len(supplier.slice) {
		return optional.Empty[T]()
	}

	curr := supplier.slice[supplier.index]
	supplier.index++
	return optional.FromValue(curr)
}

func (supplier *sliceSupplier[T]) NextRef() optional.Optional[*T] {
	if supplier.index >= len(supplier.slice) {
		return optional.Empty[*T]()
	}

	curr := &supplier.slice[supplier.index]
	supplier.index++
	return optional.FromValue(curr)
}

func Slice[T any](slice []T) Supplier[T] {
	supplier := sliceSupplier[T]{slice, 0}
	return Function[T](supplier.NextValue)
}

func SliceRefs[T any](slice []T) Supplier[*T] {
	supplier := sliceSupplier[T]{slice, 0}
	return Function[*T](supplier.NextRef)
}
