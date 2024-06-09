package misc

import "fmt"

// Pair is a generic type that
// holds two values of different types.
type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// MakePair creates a new Pair[T1, T2] with the given values.
// It is a convenient function to avoid having to specify
// the types when creating a Pair.
func MakePair[T1, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{first, second}
}

// Get returns the values as a tuple.
func (pair Pair[T1, T2]) Get() (T1, T2) { return pair.First, pair.Second }

// String returns the string representation.
func (pair Pair[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", pair.First, pair.Second)
}
