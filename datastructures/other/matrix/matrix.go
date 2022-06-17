package matrix

type Matrix[T any] struct {
	slice [][]T
}

func New[T any](height, width int) *Matrix[T] {
	arr := make([][]T, height)
	for i := 0; i < height; i++ {
		arr[i] = make([]T, width)
	}

	return &Matrix[T]{arr}
}

func (matrix *Matrix[T]) Size() (height, width int) {
	height = len(matrix.slice)

	if height > 0 {
		width = len(matrix.slice[0])
	} else {
		width = 0
	}

	return
}

func (matrix *Matrix[T]) Get(i, j int) T {
	return matrix.slice[i][j]
}

func (matrix *Matrix[T]) Set(i, j int, value T) {
	matrix.slice[i][j] = value
}

func (matrix *Matrix[T]) Clone() *Matrix[T] {
	height, width := matrix.Size()

	newMatrix := New[T](height, width)
	for i := 0; i < height; i++ {
		copy(newMatrix.slice[i], matrix.slice[i])
	}

	return newMatrix
}
