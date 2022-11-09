package matrix

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Complex | constraints.Float | constraints.Integer
}

func Add[T number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		//TODO: Improve panic type
		panic("Matrix sizes don't match")
	}

	height, width := first.Size()

	result := NewWithSize[T](height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sum := first.Get(i, j) + second.Get(i, j)
			result.Set(i, j, sum)
		}
	}

	return result
}

func Subtract[T number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		//TODO: Improve panic type
		panic("Matrix sizes don't match")
	}

	height, width := first.Size()

	result := NewWithSize[T](height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			diff := first.Get(i, j) - second.Get(i, j)
			result.Set(i, j, diff)
		}
	}

	return result
}
