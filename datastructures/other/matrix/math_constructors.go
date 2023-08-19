package matrix

import "github.com/djordje200179/extendedlibrary/misc/math"

func Zeros[T math.Number](size Size) *Matrix[T] {
	return New[T](size)
}

func Ones[T math.Number](size Size) *Matrix[T] {
	matrix := New[T](size)

	for i := 0; i < size.Elements(); i++ {
		matrix.values[i] = 1
	}

	return matrix
}

func Identity[T math.Number](size Size) *Matrix[T] {
	matrix := Zeros[T](size)

	for i := 0; i < size.Height && i < size.Width; i++ {
		matrix.Set(i, i, 1)
	}

	return matrix
}
