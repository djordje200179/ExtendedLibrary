package functions

import "github.com/djordje200179/extendedlibrary/misc/comparison"

func Increment[T integer](i T) T {
	return i + 1
}

func Ascending[T ordered](a, b T) comparison.Result {
	switch {
	case a < b:
		return comparison.FirstSmaller
	case a > b:
		return comparison.SecondSmaller
	default:
		return comparison.Equal
	}
}

func Descending[T ordered](a, b T) comparison.Result {
	switch {
	case a < b:
		return comparison.FirstBigger
	case a > b:
		return comparison.SecondBigger
	default:
		return comparison.Equal
	}
}
