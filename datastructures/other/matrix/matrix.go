package matrix

type Matrix[T any] struct {
	rows [][]T
}

func New[T any]() *Matrix[T] {
	return NewWithSize[T](0, 0)
}

func NewWithCapacity[T any](initialHeightCapacity, initialWidthCapacity int) *Matrix[T] {
	rows := make([][]T, 0, initialHeightCapacity)
	for i := 0; i < initialHeightCapacity; i++ {
		rows[i] = make([]T, 0, initialWidthCapacity)
	}

	matrix := new(Matrix[T])
	matrix.rows = rows

	return matrix
}

func NewWithSize[T any](initialHeight, initialWidth int) *Matrix[T] {
	rows := make([][]T, initialHeight)
	for i := 0; i < initialHeight; i++ {
		rows[i] = make([]T, initialWidth)
	}

	matrix := new(Matrix[T])
	matrix.rows = rows

	return matrix
}

func (matrix *Matrix[T]) Size() (height, width int) {
	height = len(matrix.rows)

	if height > 0 {
		width = len(matrix.rows[0])
	} else {
		width = 0
	}

	return
}

func (matrix *Matrix[T]) Get(i, j int) T {
	return matrix.rows[i][j]
}

func (matrix *Matrix[T]) Set(i, j int, value T) {
	matrix.rows[i][j] = value
}

func (matrix *Matrix[T]) Clone() *Matrix[T] {
	height, width := matrix.Size()

	newMatrix := NewWithSize[T](height, width)
	for i := 0; i < height; i++ {
		copy(newMatrix.rows[i], matrix.rows[i])
	}

	return newMatrix
}

func (matrix *Matrix[T]) AppendRow(row []T) {
	if _, oldWidth := matrix.Size(); len(row) != oldWidth {
		panic("runtime error: row length does not match matrix width")
	}

	matrix.rows = append(matrix.rows, row)
}

func (matrix *Matrix[T]) AppendColumn(column []T) {
	if oldHeight, _ := matrix.Size(); len(column) != oldHeight {
		panic("runtime error: column length does not match matrix height")
	}

	for i := 0; i < len(matrix.rows); i++ {
		matrix.rows[i] = append(matrix.rows[i], column[i])
	}
}
