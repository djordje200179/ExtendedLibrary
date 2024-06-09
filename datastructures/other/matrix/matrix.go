package matrix

import (
	"errors"
	"slices"
)

// Matrix is a two-dimensional dynamic array of values.
//
// The zero value is ready to use.
// Do not copy a non-zero Matrix.
type Matrix[T any] struct {
	values []T

	columns int
}

// New creates an empty Matrix with the specified Size.
func New[T any](size Size) *Matrix[T] {
	values := make([]T, size.Elements())

	matrix := &Matrix[T]{
		values:  values,
		columns: size.Width,
	}

	return matrix
}

// SliceSizeMismatchError is an error that occurs
// when the size of the given row or column slice
// doesn't match the size of the matrix.
var SliceSizeMismatchError = errors.New("slice size isn't consistent")

// NewFromSlices creates a new Matrix from
// the elements of the specified slices.
//
// The slices must have the same length,
// SliceSizeMismatchError panics occur otherwise.
func NewFromSlices[T any](values [][]T) *Matrix[T] {
	if len(values) == 0 {
		return New[T](Size{})
	}

	rows := len(values)
	columns := len(values[0])

	for _, row := range values {
		if len(row) != columns {
			panic(SliceSizeMismatchError)
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

// Size returns the Size.
func (matrix *Matrix[T]) Size() Size {
	var rows int
	if matrix.columns != 0 {
		rows = len(matrix.values) / matrix.columns
	}

	return Size{rows, matrix.columns}
}

// GetRef returns the reference to the element at the specified position.
func (matrix *Matrix[T]) GetRef(row, column int) *T {
	return &matrix.values[matrix.Size().Index(row, column)]
}

// Get returns the element at the specified position.
func (matrix *Matrix[T]) Get(row, column int) T { return *matrix.GetRef(row, column) }

// Set sets the element at the specified position.
func (matrix *Matrix[T]) Set(row, column int, value T) { *matrix.GetRef(row, column) = value }

// Clone returns a copy of the Matrix.
func (matrix *Matrix[T]) Clone() *Matrix[T] {
	return &Matrix[T]{
		values:  slices.Clone(matrix.values),
		columns: matrix.columns,
	}
}

// InsertRow inserts the given row at the given index.
//
// The row length must match the matrix width,
// SliceSizeMismatchError panics occur otherwise.
func (matrix *Matrix[T]) InsertRow(index int, row []T) {
	if len(row) != matrix.columns {
		panic(SliceSizeMismatchError)
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

// InsertColumn inserts the given column at the given index.
//
// The column length must match the matrix height,
// SliceSizeMismatchError panics occur otherwise.
func (matrix *Matrix[T]) InsertColumn(index int, column []T) {
	size := matrix.Size()

	if len(column) != size.Height {
		panic(SliceSizeMismatchError)
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

// AppendRow appends the given row to the end.
//
// The row length must match the Matrix width,
// SliceSizeMismatchError panics occur otherwise.
func (matrix *Matrix[T]) AppendRow(row []T) {
	if len(row) != matrix.columns {
		panic(SliceSizeMismatchError)
	}

	matrix.values = append(matrix.values, row...)
}

// AppendColumn appends the given column to the end.
//
// The column length must match the Matrix height,
// SliceSizeMismatchError panics occur otherwise.
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

// SizeMismatchError is an error that occurs
// when the sizes of two matrices don't match.
var SizeMismatchError = errors.New("matrix sizes don't match")

// Reshape reshapes the Matrix into the
// given size in row-major order.
//
// The new size must have the same number of elements,
// SizeMismatchError panics occur otherwise.
func (matrix *Matrix[T]) Reshape(newSize Size) {
	oldSize := matrix.Size()

	if oldSize.Elements() != newSize.Elements() {
		panic(SizeMismatchError)
	}

	matrix.columns = newSize.Width
}

// Transpose transposes the Matrix.
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
