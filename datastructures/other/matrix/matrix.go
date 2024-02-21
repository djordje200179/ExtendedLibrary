package matrix

import "fmt"

// Matrix is a two-dimensional array of values.
// The zero value is ready to use. Do not copy a non-zero Matrix.
type Matrix[T any] struct {
	values []T

	columns int
}

// New creates a new matrix with the given size.
func New[T any](size Size) *Matrix[T] {
	values := make([]T, size.Elements())

	matrix := &Matrix[T]{
		values:  values,
		columns: size.Width,
	}

	return matrix
}

// NewFromSlices creates a new matrix from the given slices.
func NewFromSlices[T any](values [][]T) *Matrix[T] {
	if len(values) == 0 {
		return New[T](Size{})
	}

	rows := len(values)
	columns := len(values[0])

	for _, row := range values {
		if len(row) != columns {
			panic("Matrix rows have different lengths")
		}
	}

	matrix := New[T](Size{rows, columns})

	for i, row := range values {
		for j, value := range row {
			matrix.Set(i, j, value)
		}
	}

	return matrix
}

// Size returns the size of the matrix.
func (matrix *Matrix[T]) Size() Size {
	var rows int
	if matrix.columns != 0 {
		rows = len(matrix.values) / matrix.columns
	}

	return Size{rows, matrix.columns}
}

// GetRef returns a reference to the value at the given position.
func (matrix *Matrix[T]) GetRef(row, column int) *T {
	index := matrix.Size().Index(row, column)
	return &matrix.values[index]
}

// Get returns the value at the given position.
func (matrix *Matrix[T]) Get(row, column int) T {
	return *matrix.GetRef(row, column)
}

// Set sets the value at the given position.
func (matrix *Matrix[T]) Set(row, column int, value T) {
	*matrix.GetRef(row, column) = value
}

// Clone returns a copy of the matrix.
func (matrix *Matrix[T]) Clone() *Matrix[T] {
	newMatrix := New[T](matrix.Size())
	copy(newMatrix.values, matrix.values)

	return newMatrix
}

// InsertRow inserts a row at the given index.
func (matrix *Matrix[T]) InsertRow(index int, row []T) {
	if len(row) != matrix.columns {
		panic("Row length does not match matrix width")
	}

	newValues := make([]T, len(matrix.values)+matrix.columns)

	oldPrevPart := matrix.values[:index*matrix.columns]
	newPrevPart := newValues[:index*matrix.columns]
	copy(newPrevPart, oldPrevPart)

	oldNextPart := matrix.values[index*matrix.columns:]
	newNextPart := newValues[(index+1)*matrix.columns:]
	copy(newNextPart, oldNextPart)

	newPart := newValues[index*matrix.columns : (index+1)*matrix.columns]
	copy(newPart, row)

	matrix.values = newValues
}

// InsertColumn inserts a column at the given index.
func (matrix *Matrix[T]) InsertColumn(index int, column []T) {
	size := matrix.Size()

	if len(column) != size.Height {
		panic("Column length does not match matrix height")
	}

	newValues := make([]T, len(matrix.values)+size.Height)

	for i := size.Height - 1; i >= 0; i-- {
		oldRow := matrix.values[i*size.Width : (i+1)*size.Width]
		newRow := newValues[i*(size.Width+1) : (i+1)*(size.Width+1)]

		copy(newRow[:index], oldRow[:index])
		copy(newRow[index+1:], oldRow[index:])
		newRow[index] = column[i]
	}

	matrix.values = newValues
	matrix.columns++
}

// AppendRow appends a row to the matrix.
func (matrix *Matrix[T]) AppendRow(row []T) {
	if len(row) != matrix.columns {
		panic("Row length does not match matrix width")
	}

	matrix.values = append(matrix.values, row...)
}

// AppendColumn appends a column to the matrix.
func (matrix *Matrix[T]) AppendColumn(column []T) {
	matrix.InsertColumn(matrix.columns, column)
}

// RemoveRow removes the row at the given index.
func (matrix *Matrix[T]) RemoveRow(index int) {
	newValues := make([]T, len(matrix.values)-matrix.columns)

	oldPrevPart := matrix.values[:index*matrix.columns]
	newPrevPart := newValues[:index*matrix.columns]
	copy(newPrevPart, oldPrevPart)

	oldNextPart := matrix.values[(index+1)*matrix.columns:]
	newNextPart := newValues[index*matrix.columns:]
	copy(newNextPart, oldNextPart)

	matrix.values = newValues
}

// RemoveColumn removes the column at the given index.
func (matrix *Matrix[T]) RemoveColumn(index int) {
	size := matrix.Size()

	newValues := make([]T, len(matrix.values)-size.Height)

	for i := range size.Height {
		oldRow := matrix.values[i*(size.Width+1) : (i+1)*(size.Width+1)]
		newRow := newValues[i*size.Width : (i+1)*size.Width]

		copy(newRow[:index], oldRow[:index])
		copy(newRow[index:], oldRow[index+1:])
	}

	matrix.values = newValues
	matrix.columns--
}

// Reshape reshapes the matrix into the given size.
// The number of elements must be the same.
// The matrix is reshaped in row-major order.
func (matrix *Matrix[T]) Reshape(newSize Size) {
	oldSize := matrix.Size()

	if oldSize.Elements() != newSize.Elements() {
		panic(fmt.Sprintf("Can't reshape matrix into %s", newSize))
	}

	matrix.columns = newSize.Width
}

// Transpose transposes the matrix.
func (matrix *Matrix[T]) Transpose() {
	oldSize := matrix.Size()
	newSize := oldSize.Transposed()

	newValues := make([]T, newSize.Elements())

	for i := range newSize.Height {
		for j := range newSize.Width {
			newIndex := newSize.Index(i, j)
			oldIndex := oldSize.Index(j, i)

			newValues[newIndex] = matrix.values[oldIndex]
		}
	}

	matrix.values = newValues
	matrix.columns = newSize.Width
}
