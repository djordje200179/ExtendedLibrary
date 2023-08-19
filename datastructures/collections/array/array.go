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

type Array[T any] []T

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
	return (*Array[T])(&slice)
}

func Collector[T any]() streams.Collector[T, *Array[T]] {
	return collections.Collector[T, *Array[T]]{New[T]()}
}

func (array *Array[T]) Size() int {
	return len(array.Slice())
}

func (array *Array[T]) Capacity() int {
	return cap(array.Slice())
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

	return &array.Slice()[index]
}

func (array *Array[T]) Get(index int) T {
	return *array.GetRef(index)
}

func (array *Array[T]) Set(index int, value T) {
	*array.GetRef(index) = value
}

func (array *Array[T]) Append(value T) {
	*array = append(array.Slice(), value)
}

func (array *Array[T]) AppendMany(values ...T) {
	*array = append(array.Slice(), values...)
}

func (array *Array[T]) Insert(index int, value T) {
	index = array.getRealIndex(index)

	newArray := make([]T, array.Size()+1)

	copy(newArray, array.Slice()[:index])
	newArray[index] = value
	copy(newArray[index+1:], array.Slice()[index:])

	*array = newArray
}

func (array *Array[T]) InsertMany(index int, values ...T) {
	index = array.getRealIndex(index)

	*array = slices.Insert(array.Slice(), index, values...)
}

func (array *Array[T]) Remove(index int) {
	index = array.getRealIndex(index)

	oldSlice := array.Slice()
	switch {
	case index == 0:
		*array = oldSlice[1:]
	case index == array.Size()-1:
		*array = oldSlice[:index]
	default:
		*array = append(oldSlice[:index], oldSlice[index+1:]...)
	}
}

func (array *Array[T]) Reserve(capacity int) {
	if capacity <= array.Capacity() {
		return
	}

	newArray := make([]T, array.Size(), capacity)
	copy(newArray, array.Slice())

	*array = newArray
}

func (array *Array[T]) Clear() {
	*array = make([]T, 0)
}

func (array *Array[T]) Reverse() {
	slices.Reverse(array.Slice())
}

func (array *Array[T]) Sort(comparator comparison.Comparator[T]) {
	slices.SortStableFunc(array.Slice(), comparator)
}

func (array *Array[T]) Join(other collections.Collection[T]) {
	switch second := other.(type) {
	case *Array[T]:
		array.AppendMany(*second...)
	default:
		for it := other.Iterator(); it.Valid(); it.Move() {
			array.Append(it.Get())
		}
	}

	other.Clear()
}

func (array *Array[T]) Clone() collections.Collection[T] {
	newArray := new(Array[T])
	*newArray = slices.Clone(*array)
	return newArray
}

func (array *Array[T]) Iterator() iterable.Iterator[T] {
	return array.ModifyingIterator()
}

func (array *Array[T]) ModifyingIterator() collections.Iterator[T] {
	return &Iterator[T]{array, 0}
}

func (array *Array[T]) Stream() streams.Stream[T] {
	supplier := suppliers.SliceValues(array.Slice())
	return streams.New(supplier)
}

func (array *Array[T]) RefsStream() streams.Stream[*T] {
	supplier := suppliers.SliceRefs(array.Slice())
	return streams.New(supplier)
}

func (array *Array[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	index := slices.IndexFunc(array.Slice(), predicate)
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
	return *array
}

func (array *Array[T]) SliceRange(from, to int) []T {
	from = array.getRealIndex(from)
	to = array.getRealIndex(to)

	return array.Slice()[from:to]
}
