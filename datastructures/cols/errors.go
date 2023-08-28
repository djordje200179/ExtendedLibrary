package cols

import "fmt"

// PanicOnIndexOutOfBounds panics with a message that the index is out of bounds
func PanicOnIndexOutOfBounds(index, length int) {
	message := fmt.Sprintf("Tried to access index %d of a collection with length %d", index, length)
	panic(message)
}
