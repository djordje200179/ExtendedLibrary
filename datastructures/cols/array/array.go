package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"slices"
)

// Array is a collection that stores elements in a contiguous memory block.
// The zero value is ready to use. Do not copy a non-zero Array.
type Array[T any] struct {
	slice []T
}

// New creates an empty Array.
func New[T any]() *Array[T] {
	return NewWithSize[T](0)
}

// NewWithSize creates an empty Array with the specified initial size.
func NewWithSize[T any](initialSize int) *Array[T] {
	return FromSlice(make([]T, initialSize))
}

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
func FromSlice[T any](slice []T) *Array[T] {
	return &Array[T]{slice}
}

// Size returns the number of elements in the Array.
func (arr *Array[T]) Size() int {
	return len(arr.slice)
}

// Capacity returns the number of elements that the Array can hold without reallocating.
func (arr *Array[T]) Capacity() int {
	return cap(arr.slice)
}

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
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (arr *Array[T]) GetRef(index int) *T {
	index = arr.getRealIndex(index)
	return &arr.slice[index]
}

// Get returns the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (arr *Array[T]) Get(index int) T {
	index = arr.getRealIndex(index)
	return arr.slice[index]
}

// Set sets the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (arr *Array[T]) Set(index int, value T) {
	index = arr.getRealIndex(index)
	arr.slice[index] = value
}

// Prepend inserts the specified element at the beginning of the Array.
func (arr *Array[T]) Prepend(value T) {
	newArray := make([]T, arr.Size()+1)

	newArray[0] = value
	copy(newArray[1:], arr.slice)

	arr.slice = newArray
}

// PrependMany inserts the specified elements at the beginning of the Array.
func (arr *Array[T]) PrependMany(values ...T) {
	newArray := make([]T, arr.Size()+len(values))

	copy(newArray, values)
	copy(newArray[len(values):], arr.slice)

	arr.slice = newArray
}

// Append appends the specified element to the end of the Array.
func (arr *Array[T]) Append(value T) {
	arr.slice = append(arr.slice, value)
}

// AppendMany appends the specified elements to the end of the Array.
func (arr *Array[T]) AppendMany(values ...T) {
	arr.slice = append(arr.slice, values...)
}

// Insert inserts the specified element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (arr *Array[T]) Insert(index int, value T) {
	index = arr.getRealIndex(index)

	newArray := make([]T, arr.Size()+1)

	copy(newArray, arr.slice[:index])
	newArray[index] = value
	copy(newArray[index+1:], arr.slice[index:])

	arr.slice = newArray
}

// InsertMany inserts the specified elements at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (arr *Array[T]) InsertMany(index int, values ...T) {
	index = arr.getRealIndex(index)

	arr.slice = slices.Insert(arr.slice, index, values...)
}

// Remove removes the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (arr *Array[T]) Remove(index int) {
	index = arr.getRealIndex(index)

	oldSlice := arr.slice
	switch {
	case index == 0:
		arr.slice = oldSlice[1:]
	case index == arr.Size()-1:
		arr.slice = oldSlice[:index]
	default:
		arr.slice = append(oldSlice[:index], oldSlice[index+1:]...)
	}
}

// Reserve reserves additional capacity for the Array.
// If additionalCapacity is negative, the function does nothing.
func (arr *Array[T]) Reserve(additionalCapacity int) {
	newCapacity := arr.Capacity() + additionalCapacity

	if newCapacity <= arr.Capacity() {
		return
	}

	newSlice := make([]T, arr.Size(), newCapacity)
	copy(newSlice, arr.slice)
	arr.slice = newSlice
}

// Shrink shrinks the capacity of the Array to match its size.
func (arr *Array[T]) Shrink() {
	arr.slice = slices.Clip(arr.slice)
}

// Clear removes all elements from the Array.
func (arr *Array[T]) Clear() {
	arr.slice = make([]T, 0)
}

// Reverse reverses the order of the elements in the Array.
func (arr *Array[T]) Reverse() {
	slices.Reverse(arr.slice)
}

// Sort sorts the elements in the Array by the specified comparator.
// The sorting algorithm is stable.
func (arr *Array[T]) Sort(comparator comparison.Comparator[T]) {
	slices.SortStableFunc(arr.slice, comparator)
}

// Join appends all elements from the other collection to the Array.
// The other collection is cleared.
func (arr *Array[T]) Join(other cols.Collection[T]) {
	switch second := other.(type) {
	case *Array[T]:
		arr.AppendMany(second.slice...)
	default:
		for it := other.Iterator(); it.Valid(); it.Move() {
			arr.Append(it.Get())
		}
	}

	other.Clear()
}

// Clone returns a copy of the Array.
func (arr *Array[T]) Clone() cols.Collection[T] {
	return &Array[T]{slices.Clone(arr.slice)}
}

// Iterator returns an iterator over the elements in the Array.
func (arr *Array[T]) Iterator() iter.Iterator[T] {
	return arr.CollectionIterator()
}

// CollectionIterator returns an iterator over the elements in the Array.
func (arr *Array[T]) CollectionIterator() cols.Iterator[T] {
	return &Iterator[T]{arr, 0}
}

// Stream streams elements of the Array.
func (arr *Array[T]) Stream(yield func(T) bool) {
	// TODO: Use slices.Stream
	for _, val := range arr.slice {
		if !yield(val) {
			break
		}
	}
}

// RefsStream streams references to the elements in the Array.
func (arr *Array[T]) RefsStream(yield func(*T) bool) {
	for i := range arr.slice {
		if !yield(&arr.slice[i]) {
			break
		}
	}
}

// FindIndex returns the index of the first element that matches the specified predicate.
func (arr *Array[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	index := slices.IndexFunc(arr.slice, predicate)
	if index == -1 {
		return 0, false
	}

	return index, true
}

// FindRef returns a reference to the first element that matches the specified predicate.
func (arr *Array[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	index, ok := arr.FindIndex(predicate)
	if !ok {
		return nil, false
	}

	return arr.GetRef(index), true
}

// Slice returns a slice of elements in the Array.
func (arr *Array[T]) Slice() []T {
	return arr.slice
}

// SliceRange returns a slice of elements in the Array in the specified range.
func (arr *Array[T]) SliceRange(from, to int) []T {
	from = arr.getRealIndex(from)
	to = arr.getRealIndex(to)

	return arr.slice[from:to]
}
