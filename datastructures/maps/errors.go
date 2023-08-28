package maps

import "fmt"

// ErrMissingKey is an error that is panicked when trying
// to access a key that is missing from a map.
type ErrMissingKey[K any] struct {
	Key K // the key that was missing
}

// Error returns the error message.
func (err ErrMissingKey[K]) Error() string {
	return fmt.Sprintf("Tried to access missing key: %v", err.Key)
}
