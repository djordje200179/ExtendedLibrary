package iterable

// An Iterator is used to iterate through a collection of elements.
// It fits well with the for-loop syntax.
type Iterator[T any] interface {
	Valid() bool // Valid returns true if the iterator points to a valid element.
	Move()       // Move moves the iterator to the next element.

	Get() T // Get returns the current element.
}

// An Iterable is a collection of elements that can be iterated through.
type Iterable[T any] interface {
	Iterator() Iterator[T] // Iterator returns an iterator for the collection.
}
