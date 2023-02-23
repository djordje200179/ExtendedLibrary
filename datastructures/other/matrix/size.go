package matrix

import "fmt"

type Size struct {
	Height, Width int
}

func (size Size) Elements() int {
	return size.Height * size.Width
}

func (size Size) Transposed() Size {
	return Size{size.Width, size.Height}
}

func (size Size) Index(row, column int) int {
	return row*size.Width + column
}

func (size Size) String() string {
	return fmt.Sprintf("(%d, %d)", size.Height, size.Width)
}
