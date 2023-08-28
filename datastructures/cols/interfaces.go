package cols

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Iterator is a special iter.Iterator that can modify the collection it iterates over.
type Iterator[T any] interface {
	iter.Iterator[T]
	// GetRef returns a reference to the current element.
	GetRef() *T
	// Set sets the current element to the given value.
	Set(value T)

	// InsertBefore inserts the given value before the current element.
	InsertBefore(value T)
	// InsertAfter inserts the given value after the current element.
	InsertAfter(value T)

	// Remove removes the current element.
	Remove()

	// Index returns the Index of the current element.
	Index() int
}

// Collection is a special iter.Iterable that represents a collection of elements.
type Collection[T any] interface {
	// Size returns the number of elements.
	Size() int

	// Get returns the element at the given index.
	Get(index int) T
	// GetRef returns a reference to the element at the given Index.
	GetRef(index int) *T
	// Set sets the element at the given Index to the given value.
	Set(index int, value T)

	// Prepend prepends the given value to the collection.
	Prepend(value T)
	// Append appends the given value to the collection.
	Append(value T)
	// Insert inserts the given value at the given index.
	Insert(index int, value T)
	// Remove removes the element at the given index.
	Remove(index int)

	// Clear clears the collection.
	Clear()
	// Reverse reverses the collection.
	Reverse()
	// Sort sorts the collection using the given comparator.
	Sort(comparator comparison.Comparator[T])
	// Join joins the collection with the given collection.
	Join(other Collection[T])

	misc.Cloner[Collection[T]]

	iter.Iterable[T]
	// CollectionIterator creates and returns a new Iterator.
	CollectionIterator() Iterator[T]
	streams.Streamer[T]
	// RefsStream returns a streams.Stream of references to the elements of the collection.
	RefsStream() streams.Stream[*T]

	// FindIndex returns the index of the first element that satisfies the given predicate.
	// If no element satisfies the predicate, 0 and false are returned.
	FindIndex(predicate predication.Predicate[T]) (int, bool)
	// FindRef returns a reference to the first element that satisfies the given predicate.
	// If no element satisfies the predicate, nil and false are returned.
	FindRef(predicate predication.Predicate[T]) (*T, bool)
}
