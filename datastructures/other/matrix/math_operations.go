package matrix

import (
	"github.com/djordje200179/extendedlibrary/misc/math"
)

// Add adds multiple matrices together and returns resulting Matrix.
// If no matrices are provided, nil is returned.
//
// If the matrices are not of the same size,
// SizeMismatchError panic occurs.
func Add[T math.Real](matrices ...*Matrix[T]) *Matrix[T] {
	if len(matrices) == 0 {
		return nil
	}

	for i := 1; i < len(matrices); i++ {
		if matrices[i].Size() != matrices[i-1].Size() {
			panic(SizeMismatchError)
		}
	}

	size := matrices[0].Size()

	result := Zeros[T](size)
	for i := range size.Elements() {
		for _, matrix := range matrices {
			result.values[i] += matrix.values[i]
		}
	}

	return result
}

// Subtract subtracts the second Matrix from the first
// and returns resulting Matrix.
//
// If the matrices are not of the same size,
// SizeMismatchError panic occurs.
func Subtract[T math.Real](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		panic(SizeMismatchError)
	}

	size := first.Size()

	result := New[T](size)

	for i := range size.Elements() {
		result.values[i] = first.values[i] - second.values[i]
	}

	return result
}

// ScalarMultiply multiplies the Matrix by
// the given scalar and returns the resulting Matrix.
func ScalarMultiply[T math.Real](matrix *Matrix[T], scalar T) *Matrix[T] {
	result := New[T](matrix.Size())

	for i := range result.Size().Elements() {
		result.values[i] = matrix.values[i] * scalar
	}

	return result
}

// Multiply multiplies two matrices together
// and returns resulting Matrix.
//
// If the width of the first matrix is not equal
// to the height of the second matrix SizeMismatchError panic occurs.
func Multiply[T math.Real](first, second *Matrix[T]) *Matrix[T] {
	if first.Size().Width != second.Size().Height {
		panic(SizeMismatchError)
	}

	resultSize := Size{
		first.Size().Height,
		second.Size().Width,
	}

	result := Zeros[T](resultSize)
	for i := range result.Size().Height {
		for j := range result.Size().Width {
			var sum T
			for k := range first.Size().Width {
				sum += first.Get(i, k) * second.Get(k, j)
			}

			result.Set(i, j, sum)
		}
	}

	return result
}

// DotMultiply multiplies element-wise multiple matrices together
// and returns resulting Matrix.
// If no matrices are provided, nil is returned.
//
// If the matrices are not of the same size,
// SizeMismatchError panic occurs.
func DotMultiply[T math.Real](matrices ...*Matrix[T]) *Matrix[T] {
	if len(matrices) == 0 {
		return nil
	}

	for i := 1; i < len(matrices); i++ {
		if matrices[i].Size() != matrices[i-1].Size() {
			panic(SizeMismatchError)
		}
	}

	size := matrices[0].Size()

	result := New[T](size)
	for i := range size.Elements() {
		result.values[i] = 1
		for _, matrix := range matrices {
			result.values[i] *= matrix.values[i]
		}
	}

	return result
}

// Negate negates the Matrix and returns the resulting Matrix.
func Negate[T math.Real](matrix *Matrix[T]) *Matrix[T] {
	return ScalarMultiply[T](matrix, -1)
}
