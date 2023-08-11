package comparison

import (
	"cmp"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

func NewByField[T any, P cmp.Ordered](getter functions.Mapper[T, P]) Comparator[T] {
	return func(first, second T) int {
		firstValue := getter(first)
		secondValue := getter(second)

		return cmp.Compare(firstValue, secondValue)
	}
}
