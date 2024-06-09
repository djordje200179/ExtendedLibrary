package matrix

import "github.com/djordje200179/extendedlibrary/misc/math"

// Zeros creates a Matrix of the given Size filled with zeros.
func Zeros[T math.Number](size Size) *Matrix[T] {
	return New[T](size)
}

// Ones creates a Matrix of the given Size filled with ones.
func Ones[T math.Number](size Size) *Matrix[T] {
	matrix := New[T](size)

	for i := range size.Elements() {
		matrix.values[i] = 1
	}

	return matrix
}

// Identity returns a Matrix of the given Size filled with
// zeros, except for the diagonal which is filled with ones.
func Identity[T math.Number](size Size) *Matrix[T] {
	matrix := Zeros[T](size)

	for i := range min(size.Height, size.Width) {
		matrix.Set(i, i, 1)
	}

	return matrix
}
