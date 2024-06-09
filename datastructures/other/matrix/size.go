package matrix

import "fmt"

// Size represents the size of a Matrix.
type Size struct {
	Height int // Height is the number of rows
	Width  int // Width is the number of columns.
}

// Elements returns the number of elements in the Matrix.
func (size Size) Elements() int { return size.Height * size.Width }

// Transposed returns the Size with the height and width swapped.
func (size Size) Transposed() Size { return Size{size.Width, size.Height} }

// Index returns the index of the element at the given position.
func (size Size) Index(row, column int) int { return row*size.Width + column }

// String returns the string representation.
func (size Size) String() string { return fmt.Sprintf("(%d, %d)", size.Height, size.Width) }
