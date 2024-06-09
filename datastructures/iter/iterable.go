package iter

import "iter"

// An Iterator is used to iterate through elements.
// It fits well with the for-loop syntax.
type Iterator[T any] interface {
	// Valid returns true if the current
	// state is valid and Get() can be called.
	Valid() bool
	// Move moves to the next element.
	Move()
	// Get returns the current element.
	Get() T
}

// An Iterable is a potentially infinite supply of elements
// that can be iterated through using an Iterator.
type Iterable[T any] interface {
	// Iterator returns a read-only Iterator over the elements.
	Iterator() Iterator[T]
}

// A FiniteIterable is an Iterable with a known number of elements.
type FiniteIterable[T any] interface {
	Iterable[T]
	// Size returns the number of elements.
	Size() int
}

// Iterate returns an iter.Seq that can be used
// with for-range syntax to iterate through the elements.
func Iterate[T any](it Iterable[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := it.Iterator(); i.Valid(); i.Move() {
			if !yield(i.Get()) {
				break
			}
		}
	}
}
