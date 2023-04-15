package matrix

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Complex | constraints.Float | constraints.Integer
}

func Add[T Number](matrices ...*Matrix[T]) *Matrix[T] {
	for i := 1; i < len(matrices); i++ {
		if matrices[i].Size() != matrices[i-1].Size() {
			panic("Matrix sizes don't match")
		}
	}

	size := matrices[0].Size()

	result := Zeros[T](size)
	for i := 0; i < size.Elements(); i++ {
		for _, matrix := range matrices {
			result.values[i] += matrix.values[i]
		}
	}

	return result
}

func Subtract[T Number](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		panic("Matrix sizes don't match")
	}

	size := first.Size()

	result := New[T](size)

	for i := 0; i < size.Elements(); i++ {
		result.values[i] = first.values[i] - second.values[i]
	}

	return result
}

func ScalarMultiply[T Number](matrix *Matrix[T], scalar T) *Matrix[T] {
	result := New[T](matrix.Size())

	for i := 0; i < result.Size().Elements(); i++ {
		result.values[i] = matrix.values[i] * scalar
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

func DotMultiply[T Number](matrices ...*Matrix[T]) *Matrix[T] {
	for i := 1; i < len(matrices); i++ {
		if matrices[i].Size() != matrices[i-1].Size() {
			panic("Matrix sizes don't match")
		}
	}

	size := matrices[0].Size()

	result := New[T](size)
	for i := 0; i < size.Elements(); i++ {
		result.values[i] = 1
		for _, matrix := range matrices {
			result.values[i] *= matrix.values[i]
		}
	}

	return result
}

func Negate[T Number](matrix *Matrix[T]) *Matrix[T] {
	return ScalarMultiply[T](matrix, -1)
}
