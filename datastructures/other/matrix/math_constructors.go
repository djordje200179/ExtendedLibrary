package matrix

import "github.com/djordje200179/extendedlibrary/misc/math"

// Zeros returns a matrix of the given size filled with zeros.
func Zeros[T math.Number](size Size) *Matrix[T] {
	return New[T](size)
}

// Ones returns a matrix of the given size filled with ones.
func Ones[T math.Number](size Size) *Matrix[T] {
	matrix := New[T](size)

	for i := range size.Elements() {
		matrix.values[i] = 1
	}

	return matrix
}

// Identity returns an identity matrix of the given size.
// Identity matrix is a square matrix with ones on the main diagonal and zeros elsewhere.
func Identity[T math.Number](size Size) *Matrix[T] {
	matrix := Zeros[T](size)

	for i := range min(size.Height, size.Width) {
		matrix.Set(i, i, 1)
	}

	return matrix
}
