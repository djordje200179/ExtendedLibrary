package cols

import "fmt"

// IndexOutOfBoundsError is an error that is panicked when trying
// to access an element of a collection with an index that is out of bounds.
type IndexOutOfBoundsError struct {
	Index  int // the index that was out of bounds
	Length int // the length of the collection
}

// Error returns the error message.
func (err IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %d out of bounds for collection of length %d", err.Index, err.Length)
}
