package maps

import "fmt"

// PanicOnMissingKey panics when a key is missing from a map.
func PanicOnMissingKey[K any](key K) {
	message := fmt.Sprintf("Tried to access missing key: %v", key)
	panic(message)
}
