package cols

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Iterator is an interface that represents an iterator over a collection
type Iterator[T any] interface {
	iterable.Iterator[T]
	GetRef() *T  // GetRef returns a pointer to the current element
	Set(value T) // Set sets the current element to the given value

	InsertBefore(value T) // InsertBefore inserts the given value before the current element
	InsertAfter(value T)  // InsertAfter inserts the given value after the current element

	Remove() // Remove removes the current element

	Index() int // Index returns the index of the current element
}

// Collection is an interface that represents a collection of elements
type Collection[T any] interface {
	Size() int // Size returns the size of the collection

	Get(index int) T        // Get returns the element at the given index
	GetRef(index int) *T    // GetRef returns a pointer to the element at the given index
	Set(index int, value T) // Set sets the element at the given index to the given value

	Prepend(value T)           // Prepend prepends the given value to the collection
	Append(value T)            // Append appends the given value to the collection
	Insert(index int, value T) // Insert inserts the given value at the given index
	Remove(index int)          // Remove removes the element at the given index

	Clear()                                   // Clear clears the collection
	Reverse()                                 // Reverse reverses the collection
	Sort(comparator comparison.Comparator[T]) // Sort sorts the collection using the given comparator
	Join(other Collection[T])                 // Join joins the collection with the given collection

	misc.Cloner[Collection[T]]

	iterable.Iterable[T]
	CollectionIterator() Iterator[T] // CollectionIterator returns an iterator over the collection
	streams.Streamer[T]
	RefsStream() streams.Stream[*T] // RefsStream returns a stream of pointers to the elements of the collection

	FindIndex(predicate predication.Predicate[T]) (int, bool) // FindIndex returns the index of the first element that satisfies the given predicate
	FindRef(predicate predication.Predicate[T]) (*T, bool)    // FindRef returns a pointer to the first element that satisfies the given predicate
}
