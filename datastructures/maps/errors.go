package maps

import "fmt"

func PanicOnMissingKey[K any](key K) {
	message := fmt.Sprintf("Tried to access missing key: %v", key)
	panic(message)
}
