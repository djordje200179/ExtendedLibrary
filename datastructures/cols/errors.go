package cols

import "fmt"

func PanicOnIndexOutOfBounds(index, length int) {
	message := fmt.Sprintf("Tried to access index %d of a collection with length %d", index, length)
	panic(message)
}
