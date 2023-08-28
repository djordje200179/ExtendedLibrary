package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
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
func (array *Array[T]) Size() int {
	return len(array.slice)
}

// Capacity returns the number of elements that the Array can hold without reallocating.
func (array *Array[T]) Capacity() int {
	return cap(array.slice)
}

func (array *Array[T]) getRealIndex(index int) int {
	size := array.Size()

	if index >= size || index < -size {
		panic(cols.ErrIndexOutOfBounds{Index: index, Length: size})
	}

	if index < 0 {
		index += size
	}

	return index
}

// GetRef returns a reference to the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array[T]) GetRef(index int) *T {
	index = array.getRealIndex(index)
	return &array.slice[index]
}

// Get returns the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array[T]) Get(index int) T {
	index = array.getRealIndex(index)
	return array.slice[index]
}

// Set sets the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array[T]) Set(index int, value T) {
	index = array.getRealIndex(index)
	array.slice[index] = value
}

// Prepend inserts the specified element at the beginning of the Array.
func (array *Array[T]) Prepend(value T) {
	newArray := make([]T, array.Size()+1)

	newArray[0] = value
	copy(newArray[1:], array.slice)

	array.slice = newArray
}

// PrependMany inserts the specified elements at the beginning of the Array.
func (array *Array[T]) PrependMany(values ...T) {
	newArray := make([]T, array.Size()+len(values))

	copy(newArray, values)
	copy(newArray[len(values):], array.slice)

	array.slice = newArray
}

// Append appends the specified element to the end of the Array.
func (array *Array[T]) Append(value T) {
	array.slice = append(array.slice, value)
}

// AppendMany appends the specified elements to the end of the Array.
func (array *Array[T]) AppendMany(values ...T) {
	array.slice = append(array.slice, values...)
}

// Insert inserts the specified element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array[T]) Insert(index int, value T) {
	index = array.getRealIndex(index)

	newArray := make([]T, array.Size()+1)

	copy(newArray, array.slice[:index])
	newArray[index] = value
	copy(newArray[index+1:], array.slice[index:])

	array.slice = newArray
}

// InsertMany inserts the specified elements at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array[T]) InsertMany(index int, values ...T) {
	index = array.getRealIndex(index)

	array.slice = slices.Insert(array.slice, index, values...)
}

// Remove removes the element at the specified index.
// Negative indices are interpreted as relative to the end of the Array.
// Panics if the index is out of bounds.
func (array *Array[T]) Remove(index int) {
	index = array.getRealIndex(index)

	oldSlice := array.slice
	switch {
	case index == 0:
		array.slice = oldSlice[1:]
	case index == array.Size()-1:
		array.slice = oldSlice[:index]
	default:
		array.slice = append(oldSlice[:index], oldSlice[index+1:]...)
	}
}

// Reserve reserves additional capacity for the Array.
// If additionalCapacity is negative, the function does nothing.
func (array *Array[T]) Reserve(additionalCapacity int) {
	newCapacity := array.Capacity() + additionalCapacity

	if newCapacity <= array.Capacity() {
		return
	}

	newSlice := make([]T, array.Size(), newCapacity)
	copy(newSlice, array.slice)
	array.slice = newSlice
}

// Shrink shrinks the capacity of the Array to match its size.
func (array *Array[T]) Shrink() {
	array.slice = slices.Clip(array.slice)
}

// Clear removes all elements from the Array.
func (array *Array[T]) Clear() {
	array.slice = make([]T, 0)
}

// Reverse reverses the order of the elements in the Array.
func (array *Array[T]) Reverse() {
	slices.Reverse(array.slice)
}

// Sort sorts the elements in the Array by the specified comparator.
// The sorting algorithm is stable.
func (array *Array[T]) Sort(comparator comparison.Comparator[T]) {
	slices.SortStableFunc(array.slice, comparator)
}

// Join appends all elements from the other collection to the Array.
// The other collection is cleared.
func (array *Array[T]) Join(other cols.Collection[T]) {
	switch second := other.(type) {
	case *Array[T]:
		array.AppendMany(second.slice...)
	default:
		for it := other.Iterator(); it.Valid(); it.Move() {
			array.Append(it.Get())
		}
	}

	other.Clear()
}

// Clone returns a copy of the Array.
func (array *Array[T]) Clone() cols.Collection[T] {
	return &Array[T]{slices.Clone(array.slice)}
}

// Iterator returns an iterator over the elements in the Array.
func (array *Array[T]) Iterator() iter.Iterator[T] {
	return array.CollectionIterator()
}

// CollectionIterator returns an iterator over the elements in the Array.
func (array *Array[T]) CollectionIterator() cols.Iterator[T] {
	return &Iterator[T]{array, 0}
}

// Stream returns a stream of the elements in the Array.
func (array *Array[T]) Stream() streams.Stream[T] {
	supplier := suppliers.SliceValues(array.slice)
	return streams.New(supplier)
}

// RefsStream returns a stream of references to the elements in the Array.
func (array *Array[T]) RefsStream() streams.Stream[*T] {
	supplier := suppliers.SliceRefs(array.slice)
	return streams.New(supplier)
}

// FindIndex returns the index of the first element that matches the specified predicate.
func (array *Array[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	index := slices.IndexFunc(array.slice, predicate)
	if index == -1 {
		return 0, false
	}

	return index, true
}

// FindRef returns a reference to the first element that matches the specified predicate.
func (array *Array[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	index, ok := array.FindIndex(predicate)
	if !ok {
		return nil, false
	}

	return array.GetRef(index), true
}

// Slice returns a slice of elements in the Array.
func (array *Array[T]) Slice() []T {
	return array.slice
}

// SliceRange returns a slice of elements in the Array in the specified range.
func (array *Array[T]) SliceRange(from, to int) []T {
	from = array.getRealIndex(from)
	to = array.getRealIndex(to)

	return array.slice[from:to]
}
