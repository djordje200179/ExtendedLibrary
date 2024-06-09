package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"slices"
)

// Array is a cols.Collection that
// stores elements in a contiguous memory block.
//
// The zero value is ready to use.
// Do not copy a non-zero Array.
type Array[T any] struct {
	slice []T
}

// New creates an empty Array.
func New[T any]() *Array[T] { return NewWithSize[T](0) }

// NewWithSize creates an empty Array with the specified initial size.
func NewWithSize[T any](initialSize int) *Array[T] { return FromSlice(make([]T, initialSize)) }

// NewWithCapacity creates an empty Array with the specified initial capacity.
func NewWithCapacity[T any](initialCapacity int) *Array[T] {
	return FromSlice(make([]T, 0, initialCapacity))
}

// NewFromIterable creates a new Array from the specified iter.Iterable.
func NewFromIterable[T any](iterable iter.Iterable[T]) *Array[T] {
	var array *Array[T]
	if finiteIterable, ok := any(iterable).(iter.FiniteIterable[T]); ok {
		array = NewWithSize[T](finiteIterable.Size())

		i := 0
		for it := iterable.Iterator(); it.Valid(); it.Move() {
			array.slice[i] = it.Get()
			i++
		}
	} else {
		array = New[T]()

		for it := iterable.Iterator(); it.Valid(); it.Move() {
			array.Append(it.Get())
		}
	}

	return array
}

// FromSlice creates a new Array from the specified slice.
func FromSlice[T any](slice []T) *Array[T] { return &Array[T]{slice} }

// FromValues creates a new Array from the specified values.
func FromValues[T any](values ...T) *Array[T] { return &Array[T]{values} }

// Size returns the number of elements.
func (arr *Array[T]) Size() int { return len(arr.slice) }

// Capacity returns the number of elements that
// can be stored without reallocating the memory.
func (arr *Array[T]) Capacity() int { return cap(arr.slice) }

func (arr *Array[T]) getRealIndex(index int) int {
	size := arr.Size()

	if index >= size || index < -size {
		panic(cols.IndexOutOfBoundsError{Index: index, Length: size})
	}

	if index < 0 {
		index += size
	}

	return index
}

// GetRef returns a reference to the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (arr *Array[T]) GetRef(index int) *T { return &arr.slice[arr.getRealIndex(index)] }

// Get returns the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (arr *Array[T]) Get(index int) T { return arr.slice[arr.getRealIndex(index)] }

// Set sets the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (arr *Array[T]) Set(index int, value T) { arr.slice[arr.getRealIndex(index)] = value }

// Prepend inserts the specified element at the beginning.
func (arr *Array[T]) Prepend(value T) { slices.Insert(arr.slice, 0, value) }

// PrependMany inserts the specified elements at the beginning.
func (arr *Array[T]) PrependMany(values ...T) { slices.Insert(arr.slice, 0, values...) }

// Append appends the specified element to the end.
func (arr *Array[T]) Append(value T) { arr.slice = append(arr.slice, value) }

// AppendMany appends the specified elements to the end.
func (arr *Array[T]) AppendMany(values ...T) { arr.slice = append(arr.slice, values...) }

// Insert inserts the specified element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (arr *Array[T]) Insert(index int, value T) {
	index = arr.getRealIndex(index)
	slices.Insert(arr.slice, index, value)
}

// InsertMany inserts the specified elements at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (arr *Array[T]) InsertMany(index int, values ...T) {
	index = arr.getRealIndex(index)
	slices.Insert(arr.slice, index, values...)
}

// Remove removes the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (arr *Array[T]) Remove(index int) {
	index = arr.getRealIndex(index)
	slices.Delete(arr.slice, index, index+1)
}

// Reserve reserves additional capacity.
//
// If additionalCapacity is negative, the function does nothing.
func (arr *Array[T]) Reserve(additionalCapacity int) {
	if additionalCapacity <= 0 {
		return
	}

	arr.slice = slices.Grow(arr.slice, additionalCapacity)
}

// Shrink shrinks the capacity to match the number of elements.
func (arr *Array[T]) Shrink() { arr.slice = slices.Clip(arr.slice) }

// Clear removes all elements.
func (arr *Array[T]) Clear() { arr.slice = nil }

// Reverse reverses the order of the elements.
func (arr *Array[T]) Reverse() { slices.Reverse(arr.slice) }

// Sort sorts the elements by the specified comparator.
//
// The sorting algorithm is a stable variant of quicksort.
func (arr *Array[T]) Sort(comparator comparison.Comparator[T]) {
	slices.SortStableFunc(arr.slice, comparator)
}

// Join moves all elements from the other cols.Collection
// to the end. The other cols.Collection becomes empty.
//
// If the other collection is an Array, only one
// memory allocation occurs and the elements are moved in bulk.
// Otherwise, the elements are moved one by one.
func (arr *Array[T]) Join(other cols.Collection[T]) {
	switch second := other.(type) {
	case *Array[T]:
		arr.AppendMany(second.slice...)
	default:
		arr.Reserve(other.Size())
		for it := other.Iterator(); it.Valid(); it.Move() {
			arr.Append(it.Get())
		}
	}

	other.Clear()
}

// Clone returns a copy of the Array.
func (arr *Array[T]) Clone() cols.Collection[T] { return &Array[T]{slices.Clone(arr.slice)} }

// Iterator returns a read-only iter.Iterator over the elements.
//
// Iteration starts from the first element.
func (arr *Array[T]) Iterator() iter.Iterator[T] { return arr.CollectionIterator() }

// CollectionIterator returns an Iterator over the elements.
// It can be used to modify the elements while iterating.
//
// Iteration starts from the first element.
func (arr *Array[T]) CollectionIterator() cols.Iterator[T] { return &Iterator[T]{arr, 0} }

// Stream streams all elements.
//
// It is not safe to modify the Array while streaming.
func (arr *Array[T]) Stream(yield func(T) bool) {
	// TODO: Use slices.Stream
	for _, val := range arr.slice {
		if !yield(val) {
			break
		}
	}
}

// Stream2 streams all elements with their indices.
//
// It is not safe to modify the Array while streaming.
func (arr *Array[T]) Stream2(yield func(int, T) bool) {
	// TODO: Use slices.Stream
	for i, val := range arr.slice {
		if !yield(i, val) {
			break
		}
	}
}

// FindIndex returns the index of the first element
// that satisfies the specified predicate.
// If no such element is found, 0 and false are returned.
func (arr *Array[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	index := slices.IndexFunc(arr.slice, predicate)
	if index == -1 {
		return 0, false
	}

	return index, true
}

// FindRef returns a reference to the first element
// that matches the specified predicate.
// If no element matches the predicate, nil and false are returned.
func (arr *Array[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	index, ok := arr.FindIndex(predicate)
	if !ok {
		return nil, false
	}

	return arr.GetRef(index), true
}

// Slice returns a slice of all elements.
func (arr *Array[T]) Slice() []T {
	return arr.slice
}

// SliceRange returns a slice of elements
// in the specified range [from, to).
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the range is out of bounds.
func (arr *Array[T]) SliceRange(from, to int) []T {
	from = arr.getRealIndex(from)
	to = arr.getRealIndex(to)

	return arr.slice[from:to]
}
