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

func ScalarMultiply[T Number](matrix *Matrix[T], scalar T) *Matrix[T] {
	result := NewWithSize[T](matrix.Size())

	for i := 0; i < result.Size().Height; i++ {
		for j := 0; j < result.Size().Width; j++ {
			result.Set(i, j, matrix.Get(i, j)*scalar)
		}
	}

	return result
}

func Multiply[T Number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size().Width != second.Size().Height {
		panic("Matrix sizes don't match")
	}

	resultSize := Size{
		first.Size().Height,
		second.Size().Width,
	}

	result := Zeros[T](resultSize)

	for i := 0; i < result.Size().Height; i++ {
		for j := 0; j < result.Size().Width; j++ {
			var sum T
			for k := 0; k < first.Size().Width; k++ {
				sum += first.Get(i, k) * second.Get(k, j)
			}

			result.Set(i, j, sum)
		}
	}

	return result
}

func DotMultiply[T Number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		panic("Matrix sizes don't match")
	}

	result := NewWithSize[T](first.Size())

	for i := 0; i < result.Size().Elements(); i++ {
		result.values[i] = first.values[i] * second.values[i]
	}

	return result
}
