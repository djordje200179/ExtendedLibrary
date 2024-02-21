package matrix

import "github.com/djordje200179/extendedlibrary/misc/math"

// Add adds multiple matrices together and returns the result.
// If the matrices are not of the same size, it panics.
// If no matrices are provided, it returns nil.
func Add[T math.Real](matrices ...*Matrix[T]) *Matrix[T] {
	if len(matrices) == 0 {
		return nil
	}

	for i := 1; i < len(matrices); i++ {
		if matrices[i].Size() != matrices[i-1].Size() {
			panic("Matrix sizes don't match")
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

// Subtract subtracts the second matrix from the first and returns the result.
// If the matrices are not of the same size, it panics.
func Subtract[T math.Real](first, second *Matrix[T]) *Matrix[T] {
	if first.Size() != second.Size() {
		panic("Matrix sizes don't match")
	}

	size := first.Size()

	result := New[T](size)

	for i := range size.Elements() {
		result.values[i] = first.values[i] - second.values[i]
	}

	return result
}

// ScalarMultiply multiplies the matrix with a scalar and returns the result.
func ScalarMultiply[T math.Real](matrix *Matrix[T], scalar T) *Matrix[T] {
	result := New[T](matrix.Size())

	for i := range result.Size().Elements() {
		result.values[i] = matrix.values[i] * scalar
	}

	return result
}

// Multiply multiplies two matrices together and returns the result.
func Multiply[T math.Real](first, second *Matrix[T]) *Matrix[T] {
	if first.Size().Width != second.Size().Height {
		panic("Matrix sizes don't match")
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

// DotMultiply multiplies multiple matrices together and returns the result.
// This function is equivalent to multiplying the matrices element by element.
// If the matrices are not of the same size, it panics.
// If no matrices are provided, it returns nil.
func DotMultiply[T math.Real](matrices ...*Matrix[T]) *Matrix[T] {
	if len(matrices) == 0 {
		return nil
	}

	for i := 1; i < len(matrices); i++ {
		if matrices[i].Size() != matrices[i-1].Size() {
			panic("Matrix sizes don't match")
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

// Negate negates the matrix and returns the result.
func Negate[T math.Real](matrix *Matrix[T]) *Matrix[T] {
	return ScalarMultiply[T](matrix, -1)
}
