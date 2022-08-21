package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type Function[T any] functions.EmptyGenerator[optional.Optional[T]]

func (supplier Function[T]) Supply() optional.Optional[T] { return supplier() }

func Infinite[T any](generator functions.EmptyGenerator[T]) Supplier[T] {
	return Function[T](func() optional.Optional[T] {
		return optional.FromValue[T](generator())
	})
}

func Range(lower, upper int) Supplier[int] {
	return Function[int](func() optional.Optional[int] {
		if lower < upper {
			curr := lower
			lower++
			return optional.FromValue(curr)
		} else {
			return optional.Empty[int]()
		}
	})
}
