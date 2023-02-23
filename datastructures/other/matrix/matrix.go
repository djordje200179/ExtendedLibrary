package matrix

import "fmt"

type Matrix[T any] struct {
	values []T

	columns int
}

func New[T any]() *Matrix[T] {
	return NewWithSize[T](Size{0, 0})
}

func NewWithSize[T any](size Size) *Matrix[T] {
	values := make([]T, size.Elements())

	matrix := &Matrix[T]{
		values:  values,
		columns: size.Width,
	}

	return matrix
}

func (matrix *Matrix[T]) Size() Size {
	rows := len(matrix.values) / matrix.columns

	return Size{rows, matrix.columns}
}

func (matrix *Matrix[T]) GetRef(row, column int) *T {
	index := matrix.Size().Index(row, column)
	return &matrix.values[index]
}

func (matrix *Matrix[T]) Get(row, column int) T {
	return *matrix.GetRef(row, column)
}

func (matrix *Matrix[T]) Set(row, column int, value T) {
	*matrix.GetRef(row, column) = value
}

func (matrix *Matrix[T]) Clone() *Matrix[T] {
	newMatrix := NewWithSize[T](matrix.Size())
	copy(newMatrix.values, matrix.values)

	return newMatrix
}

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

func (matrix *Matrix[T]) AppendRow(row []T) {
	if len(row) != matrix.columns {
		panic("Row length does not match matrix width")
	}

	matrix.values = append(matrix.values, row...)
}

func (matrix *Matrix[T]) AppendColumn(column []T) {
	matrix.InsertColumn(matrix.columns, column)
}

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

func (matrix *Matrix[T]) RemoveColumn(index int) {
	size := matrix.Size()

	newValues := make([]T, len(matrix.values)-size.Height)

	for i := 0; i < size.Height; i++ {
		oldRow := matrix.values[i*(size.Width+1) : (i+1)*(size.Width+1)]
		newRow := newValues[i*size.Width : (i+1)*size.Width]

		copy(newRow[:index], oldRow[:index])
		copy(newRow[index:], oldRow[index+1:])
	}

	matrix.values = newValues
	matrix.columns--
}

func (matrix *Matrix[T]) Reshape(newSize Size) {
	oldSize := matrix.Size()

	if oldSize.Elements() != newSize.Elements() {
		panic(fmt.Sprintf("Can't reshape matrix into %s", newSize))
	}

	matrix.columns = newSize.Width
}

func (matrix *Matrix[T]) Transpose() {
	oldSize := matrix.Size()
	newSize := oldSize.Transposed()

	newValues := make([]T, newSize.Elements())

	for i := 0; i < newSize.Height; i++ {
		for j := 0; j < newSize.Width; j++ {
			newIndex := newSize.Index(i, j)
			oldIndex := oldSize.Index(j, i)

			newValues[newIndex] = matrix.values[oldIndex]
		}
	}

	matrix.values = newValues
	matrix.columns = newSize.Width
}
