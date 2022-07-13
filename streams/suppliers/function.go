package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type functionSupplier[T any] functions.EmptyGenerator[optional.Optional[T]]

func FromFiniteGenerator[T any](generator functions.EmptyGenerator[optional.Optional[T]]) Supplier[T] {
	return functionSupplier[T](generator)
}

func FromInfiniteGenerator[T any](generator functions.EmptyGenerator[T]) Supplier[T] {
	return FromFiniteGenerator[T](func() optional.Optional[T] { return optional.FromValue[T](generator()) })
}

func FromRange(lower, upper int) Supplier[int] {
	return FromFiniteGenerator(func() optional.Optional[int] {
		if lower < upper {
			curr := lower
			lower++
			return optional.FromValue(curr)
		} else {
			return optional.Empty[int]()
		}
	})
}

func (supplier functionSupplier[T]) Supply() optional.Optional[T] { return supplier() }
