package matrix

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Complex | constraints.Float | constraints.Integer
}

func Add[T Number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		panic("Matrix sizes don't match")
	}

	size := first.Size()

	result := NewWithSize[T](size)

	for i := 0; i < size.Height; i++ {
		for j := 0; j < size.Width; j++ {
			sum := first.Get(i, j) + second.Get(i, j)
			result.Set(i, j, sum)
		}
	}

	return result
}

func Subtract[T Number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		panic("Matrix sizes don't match")
	}

	size := first.Size()

	result := NewWithSize[T](size)

	for i := 0; i < size.Height; i++ {
		for j := 0; j < size.Width; j++ {
			diff := first.Get(i, j) - second.Get(i, j)
			result.Set(i, j, diff)
		}
	}

	return result
}
