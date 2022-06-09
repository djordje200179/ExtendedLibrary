package functions

import "github.com/djordje200179/GoExtendedLibrary/misc/comparison"

func Increment(i int) int {
	return i + 1
}

func NumsAscending(a, b int) comparison.Result {
	return comparison.FromInt(b - a)
}

func NumsDescending(a, b int) comparison.Result {
	return comparison.FromInt(a - b)
}
