package matrix

import "fmt"

// Size represents the size of a matrix.
type Size struct {
	Height, Width int // Height is the number of rows, Width is the number of columns.
}

// Elements returns the number of elements in the matrix of the given size.
func (size Size) Elements() int {
	return size.Height * size.Width
}

// Transposed returns the size of a transposed matrix.
func (size Size) Transposed() Size {
	return Size{size.Width, size.Height}
}

// Index returns the index of the element at the given row and column.
func (size Size) Index(row, column int) int {
	return row*size.Width + column
}

// String returns a string representation of the size.
func (size Size) String() string {
	return fmt.Sprintf("(%d, %d)", size.Height, size.Width)
}
