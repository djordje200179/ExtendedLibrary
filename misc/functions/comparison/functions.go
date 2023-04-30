package comparison

import "golang.org/x/exp/constraints"

func Compare[T constraints.Ordered](a, b T) Result {
	switch {
	case a < b:
		return FirstSmaller
	case a > b:
		return SecondSmaller
	default:
		return Equal
	}
}

func ReverseCompare[T constraints.Ordered](a, b T) Result {
	switch {
	case a < b:
		return FirstBigger
	case a > b:
		return SecondBigger
	default:
		return Equal
	}
}
