package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"slices"
)

type Array[T any] struct {
	slice []T
}

func New[T any]() *Array[T] {
	return NewWithSize[T](0)
}

func NewWithSize[T any](initialSize int) *Array[T] {
	return FromSlice(make([]T, initialSize))
}

func NewWithCapacity[T any](initialCapacity int) *Array[T] {
	return FromSlice(make([]T, 0, initialCapacity))
}

func FromSlice[T any](slice []T) *Array[T] {
	return &Array[T]{slice}
}

func Collector[T any]() streams.Collector[T, *Array[T]] {
	return collections.Collector[T, *Array[T]]{New[T]()}
}

func (array *Array[T]) Size() int {
	return len(array.slice)
}

func (array *Array[T]) Capacity() int {
	return cap(array.slice)
}

func (array *Array[T]) getRealIndex(index int) int {
	size := array.Size()

	if index >= size || index < -size {
		collections.PanicOnIndexOutOfBounds(index, size)
	}

	if index < 0 {
		index += size
	}

	return index
}

func (array *Array[T]) GetRef(index int) *T {
	index = array.getRealIndex(index)

	return &array.slice[index]
}

func (array *Array[T]) Get(index int) T {
	return *array.GetRef(index)
}

func (array *Array[T]) Set(index int, value T) {
	*array.GetRef(index) = value
}

func (array *Array[T]) Append(value T) {
	array.slice = append(array.slice, value)
}

func (array *Array[T]) AppendMany(values ...T) {
	array.slice = append(array.slice, values...)
}

func (array *Array[T]) Insert(index int, value T) {
	index = array.getRealIndex(index)

	newArray := make([]T, array.Size()+1)

	copy(newArray, array.slice[:index])
	newArray[index] = value
	copy(newArray[index+1:], array.slice[index:])

	array.slice = newArray
}

func (array *Array[T]) InsertMany(index int, values ...T) {
	index = array.getRealIndex(index)

	array.slice = slices.Insert(array.slice, index, values...)
}

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

func (array *Array[T]) Reserve(capacity int) {
	if capacity <= array.Capacity() {
		return
	}

	newArray := make([]T, array.Size(), capacity)
	copy(newArray, array.slice)

	array.slice = newArray
}

func (array *Array[T]) Clear() {
	array.slice = make([]T, 0)
}

func (array *Array[T]) Reverse() {
	slices.Reverse(array.slice)
}

func (array *Array[T]) Sort(comparator comparison.Comparator[T]) {
	slices.SortStableFunc(array.slice, comparator)
}

func (array *Array[T]) Join(other collections.Collection[T]) {
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

func (array *Array[T]) Clone() collections.Collection[T] {
	return &Array[T]{slices.Clone(array.slice)}
}

func (array *Array[T]) Iterator() iterable.Iterator[T] {
	return array.CollectionIterator()
}

func (array *Array[T]) CollectionIterator() collections.Iterator[T] {
	return &Iterator[T]{array, 0}
}

func (array *Array[T]) Stream() streams.Stream[T] {
	supplier := suppliers.SliceValues(array.slice)
	return streams.New(supplier)
}

func (array *Array[T]) RefsStream() streams.Stream[*T] {
	supplier := suppliers.SliceRefs(array.slice)
	return streams.New(supplier)
}

func (array *Array[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	index := slices.IndexFunc(array.slice, predicate)
	if index == -1 {
		return 0, false
	}

	return index, true
}

func (array *Array[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	index, ok := array.FindIndex(predicate)
	if !ok {
		return nil, false
	}

	return array.GetRef(index), true
}

func (array *Array[T]) Slice() []T {
	return array.slice
}

func (array *Array[T]) SliceRange(from, to int) []T {
	from = array.getRealIndex(from)
	to = array.getRealIndex(to)

	return array.slice[from:to]
}
