package matrix

type Matrix[T any] struct {
	rows [][]T
}

type Size struct {
	Height, Width int
}

func New[T any]() *Matrix[T] {
	return NewWithSize[T](Size{0, 0})
}

func NewWithSize[T any](size Size) *Matrix[T] {
	rows := make([][]T, size.Height)
	for i := 0; i < size.Height; i++ {
		rows[i] = make([]T, size.Width)
	}

	matrix := &Matrix[T]{
		rows: rows,
	}

	return matrix
}

func (matrix *Matrix[T]) Size() Size {
	height := len(matrix.rows)

	var width int
	if height > 0 {
		width = len(matrix.rows[0])
	} else {
		width = 0
	}

	return Size{height, width}
}

func (matrix *Matrix[T]) GetRef(row, column int) *T {
	return &matrix.rows[row][column]
}

func (matrix *Matrix[T]) Get(row, column int) T {
	return matrix.rows[row][column]
}

func (matrix *Matrix[T]) Set(row, column int, value T) {
	matrix.rows[row][column] = value
}

func (matrix *Matrix[T]) Clone() *Matrix[T] {
	size := matrix.Size()

	newMatrix := NewWithSize[T](size)
	for i := 0; i < size.Height; i++ {
		copy(newMatrix.rows[i], matrix.rows[i])
	}

	return newMatrix
}

func (matrix *Matrix[T]) InsertRow(index int, row []T) {

}

func (matrix *Matrix[T]) InsertColumn(index int, column []T) {

}

func (matrix *Matrix[T]) AppendRow(row []T) {
	if len(row) != matrix.Size().Width {
		panic("runtime error: row length does not match matrix width")
	}

	matrix.rows = append(matrix.rows, row)
}

func (matrix *Matrix[T]) AppendColumn(column []T) {
	if len(column) != matrix.Size().Height {
		panic("runtime error: column length does not match matrix height")
	}

	for i := 0; i < len(matrix.rows); i++ {
		matrix.rows[i] = append(matrix.rows[i], column[i])
	}
}

func (matrix *Matrix[T]) Reshape(newSize Size) {
	oldSize := matrix.Size()

	if oldSize.Height*oldSize.Width != newSize.Height*newSize.Width {
		panic("runtime error: can't reshape matrix into new size")
	}

	newRows := NewWithSize[T](newSize).rows

	currRow, currColumn := 0, 0
	for i := 0; i < oldSize.Width; i++ {
		for j := 0; j < oldSize.Height; j++ {
			newRows[currRow][currColumn] = matrix.rows[i][j]

			currColumn++
			if currColumn == newSize.Width {
				currRow++
				currColumn = 0
			}
		}
	}

	matrix.rows = newRows
}

func (matrix *Matrix[T]) Transpose() {
	oldSize := matrix.Size()
	newSize := Size{oldSize.Width, oldSize.Height}

	newRows := NewWithSize[T](newSize).rows

	for i := 0; i < oldSize.Height; i++ {
		for j := 0; j < oldSize.Width; j++ {
			newRows[j][i] = matrix.rows[i][j]
		}
	}

	matrix.rows = newRows
}
