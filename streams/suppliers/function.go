package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type FunctionSupplier[T any] func() optional.Optional[T]

func (supplier FunctionSupplier[T]) Supply() optional.Optional[T] {
	return supplier()
}

func Infinite[T any](generator func() T) Supplier[T] {
	return FunctionSupplier[T](func() optional.Optional[T] {
		return optional.FromValue[T](generator())
	})
}

func Range(lower, upper int) Supplier[int] {
	return FunctionSupplier[int](func() optional.Optional[int] {
		if lower < upper {
			curr := lower
			lower++
			return optional.FromValue(curr)
		} else {
			return optional.Empty[int]()
		}
	})
}
