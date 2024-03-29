package iter

// An Iterator is used to iterate through elements.
// It fits well with the for-loop syntax.
type Iterator[T any] interface {
	// Valid returns true if the end has not been reached.
	Valid() bool
	// Move fetches the next element.
	Move()
	// Get returns the current element.
	Get() T
}

// An Iterable is a finite or infinite collection of elements.
// It can be iterated through using an Iterator.
type Iterable[T any] interface {
	// Iterator creates and returns a new Iterator.
	Iterator() Iterator[T]
}

// A FiniteIterable is a finite collection of elements.
// It can be iterated through using an Iterator.
// Useful for passing to data structures constructors.
type FiniteIterable[T any] interface {
	Iterable[T]
	// Size returns the number of elements in the collection.
	Size() int
}
