package matrix

func Zeros[T Number](size Size) *Matrix[T] {
	return NewWithSize[T](size)
}

func Ones[T Number](size Size) *Matrix[T] {
	matrix := NewWithSize[T](size)

	for i := 0; i < size.Elements(); i++ {
		matrix.values[i] = 1
	}

	return matrix
}

func Identity[T Number](size Size) *Matrix[T] {
	matrix := Zeros[T](size)

	for i := 0; i < size.Height && i < size.Width; i++ {
		matrix.Set(i, i, 1)
	}

	return matrix
}
