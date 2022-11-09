package comparison

import "golang.org/x/exp/constraints"

func Ascending[T constraints.Ordered](a, b T) Result {
	switch {
	case a < b:
		return FirstSmaller
	case a > b:
		return SecondSmaller
	default:
		return Equal
	}
}

func Descending[T constraints.Ordered](a, b T) Result {
	switch {
	case a < b:
		return FirstBigger
	case a > b:
		return SecondBigger
	default:
		return Equal
	}
}
