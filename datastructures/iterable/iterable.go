package iterable

// An Iterator fits well with iteration through for loop.
type Iterator[T any] interface {
	Valid() bool // Valid returns false if the iterator is exhausted.
	Move()       // Move moves the iterator to the next element.

	Get() T // Get returns the current element.
}

// An Iterable is a collection of elements that can be iterated through.
type Iterable[T any] interface {
	Iterator() Iterator[T] // Iterator returns an iterator for the collection.
}
