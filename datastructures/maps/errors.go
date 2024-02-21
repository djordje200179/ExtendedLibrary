package maps

import "fmt"

// MissingKeyError is an error that is panicked when trying
// to access a key that is missing from a map.
type MissingKeyError[K any] struct {
	Key K // the key that was missing
}

// Error returns the error message.
func (err MissingKeyError[K]) Error() string {
	return fmt.Sprintf("key %v is missing from map", err.Key)
}
