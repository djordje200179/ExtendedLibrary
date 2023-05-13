package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
	"golang.org/x/exp/constraints"
)

func Infinite[T any](generator func() T) streams.Supplier[T] {
	return func() optional.Optional[T] {
		return optional.FromValue(generator())
	}
}

func Range[T constraints.Integer | constraints.Float](lower, upper T) streams.Supplier[T] {
	return RangeWithIncrement(lower, upper, 1)
}

func RangeWithIncrement[T constraints.Integer | constraints.Float](lower, upper, increment T) streams.Supplier[T] {
	return func() optional.Optional[T] {
		if lower < upper {
			old := lower
			lower += increment
			return optional.FromValue(old)
		}

		return optional.Empty[T]()
	}
}
