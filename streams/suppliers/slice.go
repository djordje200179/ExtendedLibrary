package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type sliceSupplier[T any] struct {
	slice []T
	index int
}

func FromSlice[T any](slice []T) Supplier[T]    { return &sliceSupplier[T]{slice, 0} }
func FromValues[T any](values ...T) Supplier[T] { return FromSlice(values) }

func (supplier *sliceSupplier[T]) Supply() optional.Optional[T] {
	if supplier.index < len(supplier.slice) {
		curr := supplier.slice[supplier.index]
		supplier.index++
		return optional.FromValue(curr)
	} else {
		return optional.Empty[T]()
	}
}
