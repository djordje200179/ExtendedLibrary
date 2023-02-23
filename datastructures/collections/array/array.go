package array

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
	"sort"
)

type Array[T any] []T

func New[T any]() *Array[T] {
	return NewWithSize[T](0)
}

func NewWithSize[T any](initialSize int) *Array[T] {
	return NewFromSlice(make([]T, initialSize))
}

func NewWithCapacity[T any](initialCapacity int) *Array[T] {
	return NewFromSlice(make([]T, 0, initialCapacity))
}

func NewFromSlice[T any](slice []T) *Array[T] {
	return (*Array[T])(&slice)
}

func Collector[T any]() streams.Collector[T, collections.Collection[T]] {
	return collections.Collector[T]{
		Collection: New[T](),
	}
}

func (array *Array[T]) Size() int {
	return len(array.Slice())
}

func (array *Array[T]) getRealIndex(index int) int {
	size := array.Size()

	if index >= size || index < -size {
		panic(fmt.Sprintf("Index out of bounds: %d", index))
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

func (array *Array[T]) Append(values ...T) {
	*array = append(array.Slice(), values...)
}

func (array *Array[T]) Insert(index int, values ...T) {
	index = array.getRealIndex(index)

	*array = append((*array)[:index], (*array)[index:]...)
	copy((*array)[index:], values)
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

func (array *Array[T]) Sort(comparator functions.Comparator[T]) {
	sort.SliceStable(array.Slice(), func(i, j int) bool {
		return comparator(array.Slice()[i], array.Slice()[j]) == comparison.FirstSmaller
	})
}

func (array *Array[T]) Join(other collections.Collection[T]) {
	switch second := other.(type) {
	case *Array[T]:
		array.Append(*second...)
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
	return streams.FromSlice(array.Slice())
}

func (array *Array[T]) RefStream() streams.Stream[*T] {
	return streams.FromSliceRefs(array.Slice())
}

func (array *Array[T]) Slice() []T {
	return *array
}
