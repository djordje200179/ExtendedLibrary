package matrix

type Matrix[T any] [][]T

func New[T any](height, width int) *Matrix[T] {
	arr := make([][]T, height)
	for i := 0; i < height; i++ {
		arr[i] = make([]T, width)
	}

	return (*Matrix[T])(&arr)
}

func (matrix *Matrix[T]) Size() (int, int) {
	return len(*matrix), len((*matrix)[0])
}

func (matrix *Matrix[T]) Get(i, j int) T {
	return (*matrix)[i][j]
}

func (matrix *Matrix[T]) Set(i, j int, value T) {
	(*matrix)[i][j] = value
}

func (matrix *Matrix[T]) Clone() *Matrix[T] {
	height, width := matrix.Size()

	newMatrix := New[T](height, width)
	for i := 0; i < height; i++ {
		copy((*newMatrix)[i], (*matrix)[i])
	}

	return newMatrix
}
