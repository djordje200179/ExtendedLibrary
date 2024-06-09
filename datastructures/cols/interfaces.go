package cols

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

// Iterator is a special iter.Iterator
// that can modify the Collection it iterates over.
type Iterator[T any] interface {
	iter.Iterator[T]
	// GetRef returns the reference to the current element.
	GetRef() *T
	// Set sets the current element.
	Set(value T)

	// InsertBefore inserts the specified element
	// before the current element.
	InsertBefore(value T)
	// InsertAfter inserts the specified element
	// after the current element.
	InsertAfter(value T)

	// Remove removes the current element.
	Remove()
}

// Collection is a special iter.Iterable
// that represents a collection of elements.
type Collection[T any] interface {
	// Size returns the number of elements.
	Size() int

	// Get returns the element at the specified index.
	//
	// Negative indices are interpreted as relative to the end.
	// Panic occurs if the index is out of bounds.
	Get(index int) T
	// GetRef returns a reference to the element at the specified index.
	//
	// Negative indices are interpreted as relative to the end.
	// Panic occurs if the index is out of bounds.
	GetRef(index int) *T
	// Set sets the element at the specified index.
	//
	// Negative indices are interpreted as relative to the end.
	// Panic occurs if the index is out of bounds.
	Set(index int, value T)

	// Prepend inserts the specified element at the beginning.
	Prepend(value T)
	// Append appends the specified element to the end.
	Append(value T)
	// Insert inserts the specified element at the specified index.
	//
	// Negative indices are interpreted as relative to the end.
	// Panic occurs if the index is out of bounds.
	Insert(index int, value T)
	// Remove removes the element at the specified index.
	//
	// Negative indices are interpreted as relative to the end.
	// Panic occurs if the index is out of bounds.
	Remove(index int)

	// Clear removes all elements.
	Clear()
	// Reverse reverses the order of the elements.
	Reverse()
	// Sort sorts the elements by the specified comparator.
	Sort(comparator comparison.Comparator[T])
	// Join moves all elements from the other cols.Collection
	// to the end. The other cols.Collection becomes empty.
	Join(other Collection[T])

	misc.Cloner[Collection[T]]

	iter.Iterable[T]
	// CollectionIterator returns an Iterator over the elements.
	// It can be used to modify the elements while iterating.
	//
	// Iteration starts from the first element.
	CollectionIterator() Iterator[T]
	// Stream streams all elements.
	Stream(yield func(T) bool)
	// Stream2 streams all elements with their indices.
	Stream2(yield func(int, T) bool)

	// FindIndex returns the index of the first element
	// that satisfies the specified predicate.
	// If no such element is found, 0 and false are returned.
	FindIndex(predicate predication.Predicate[T]) (int, bool)
	// FindRef returns a reference to the first element
	// that matches the specified predicate.
	// If no element matches the predicate, nil and false are returned.
	FindRef(predicate predication.Predicate[T]) (*T, bool)
}
