package cols

import "fmt"

// ErrIndexOutOfBounds is an error that is panicked when trying
// to access an element of a collection with an index that is out of bounds.
type ErrIndexOutOfBounds struct {
	Index  int // the index that was out of bounds
	Length int // the length of the collection
}

// Error returns the error message.
func (err ErrIndexOutOfBounds) Error() string {
	return fmt.Sprintf("Tried to access index %d of a collection with length %d", err.Index, err.Length)
}
