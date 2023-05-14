package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"golang.org/x/exp/slices"
	"sort"
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

func Collector[T any]() streams.Collector[T, collections.Collection[T]] {
	return collections.Collector[T]{New[T]()}
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

	newArray := make([]T, array.Size()+len(values))

	copy(newArray, array.Slice()[:index])
	copy(newArray[index:], values)
	copy(newArray[index+len(values):], array.Slice()[index:])

	*array = newArray
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

func (array *Array[T]) Clear() {
	*array = make([]T, 0)
}

func (array *Array[T]) Reverse() {
	n := len(array.Slice())
	for i := 0; i < n/2; i++ {
		array.Slice()[i], array.Slice()[n-1-i] = array.Slice()[n-1-i], array.Slice()[i]
	}
}

func (array *Array[T]) Swap(index1, index2 int) {
	index1 = array.getRealIndex(index1)
	index2 = array.getRealIndex(index2)

	array.Slice()[index1], array.Slice()[index2] = array.Slice()[index2], array.Slice()[index1]
}

func (array *Array[T]) Sort(comparator comparison.Comparator[T]) {
	sort.SliceStable(array.Slice(), func(i, j int) bool {
		return comparator(array.Slice()[i], array.Slice()[j]) == comparison.FirstSmaller
	})
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
	cloned := NewWithSize[T](array.Size())
	copy(cloned.Slice(), array.Slice())

	return cloned
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

func (array *Array[T]) Slice() []T {
	return *array
}

func (array *Array[T]) SliceRange(from, to int) []T {
	from = array.getRealIndex(from)
	to = array.getRealIndex(to)

	return array.Slice()[from:to]
}
