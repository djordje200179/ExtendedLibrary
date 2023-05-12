package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

func Infinite[T any](generator func() T) streams.Supplier[T] {
	return func() optional.Optional[T] {
		return optional.FromValue(generator())
	}
}

func Range(lower, upper int) streams.Supplier[int] {
	return func() optional.Optional[int] {
		if lower < upper {
			lower++
			return optional.FromValue(lower - 1)
		}

		return optional.Empty[int]()
	}
}
