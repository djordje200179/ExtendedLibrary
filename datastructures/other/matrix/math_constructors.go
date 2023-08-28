package matrix

import "github.com/djordje200179/extendedlibrary/misc/math"

// Zeros returns a matrix of the given size filled with zeros.
func Zeros[T math.Number](size Size) *Matrix[T] {
	return New[T](size)
}

// Ones returns a matrix of the given size filled with ones.
func Ones[T math.Number](size Size) *Matrix[T] {
	matrix := New[T](size)

	for i := 0; i < size.Elements(); i++ {
		matrix.values[i] = 1
	}

	return matrix
}

// Identity returns an identity matrix of the given size.
// Identity matrix is a square matrix with ones on the main diagonal and zeros elsewhere.
func Identity[T math.Number](size Size) *Matrix[T] {
	matrix := Zeros[T](size)

	for i := 0; i < size.Height && i < size.Width; i++ {
		matrix.Set(i, i, 1)
	}

	return matrix
}
