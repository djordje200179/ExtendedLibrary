package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/math"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

func Infinite[T any](generator func() T) streams.Supplier[T] {
	return func() optional.Optional[T] {
		return optional.FromValue(generator())
	}
}

func Range[T math.Real](lower, upper T) streams.Supplier[T] {
	return RangeWithIncrement(lower, upper, 1)
}

func RangeWithIncrement[T math.Real](lower, upper, increment T) streams.Supplier[T] {
	return func() optional.Optional[T] {
		if lower < upper {
			old := lower
			lower += increment
			return optional.FromValue(old)
		}

		return optional.Empty[T]()
	}
}
