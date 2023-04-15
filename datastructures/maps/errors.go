package maps

import "fmt"

func PanicOnMissingKey[K comparable](key K) {
	message := fmt.Sprintf("Tried to access missing key: %v", key)
	panic(message)
}
