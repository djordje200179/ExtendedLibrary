package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sort"
)

type Array[T any] []T

func New[T any]() *Array[T] {
	return NewWithCapacity[T](0)
}

func NewWithSize[T any](initialSize int) *Array[T] {
	slice := make([]T, initialSize)
	return (*Array[T])(&slice)
}

func NewWithCapacity[T any](initialCapacity int) *Array[T] {
	slice := make([]T, 0, initialCapacity)
	return (*Array[T])(&slice)
}

func Collector[T any]() streams.Collector[T, sequences.Sequence[T]] {
	return sequences.Collector[T](New[T]())
}

func (array *Array[T]) Size() int {
	return len(array.Slice())
}

func (array *Array[T]) Get(i int) T {
	return array.Slice()[i]
}

func (array *Array[T]) Set(i int, value T) {
	array.Slice()[i] = value
}

func (array *Array[T]) Append(value T) {
	*array = append(array.Slice(), value)
}

func (array *Array[T]) AppendMany(values ...T) {
	*array = append(array.Slice(), values...)
}

func (array *Array[T]) Insert(index int, value T) {
	oldSlice := array.Slice()
	*array = append(oldSlice[:index+1], oldSlice[index:]...)
	array.Slice()[index] = value
}

func (array *Array[T]) Remove(index int) {
	oldSlice := array.Slice()
	*array = append(oldSlice[:index], oldSlice[index+1:]...)
}

func (array *Array[T]) Clear() {
	*array = nil
}

func (array *Array[T]) Reverse() {
	n := len(array.Slice())
	for i := 0; i < n/2; i++ {
		array.Slice()[i], array.Slice()[n-1-i] = array.Slice()[n-1-i], array.Slice()[i]
	}
}

func (array *Array[T]) Sort(comparator functions.Comparator[T]) {
	sort.SliceStable(array.Slice(), func(i, j int) bool {
		return comparator(array.Slice()[i], array.Slice()[j]) == comparison.FirstSmaller
	})
}

func (array *Array[T]) Join(other sequences.Sequence[T]) {
	switch second := other.(type) {
	case *Array[T]:
		array.AppendMany(second.Slice()...)
	default:
		for it := other.Iterator(); it.Valid(); it.Move() {
			array.Append(it.Get())
		}
	}
}

func (array *Array[T]) Clone() sequences.Sequence[T] {
	cloned := NewWithSize[T](array.Size())
	copy(cloned.Slice(), array.Slice())

	return cloned
}

func (array *Array[T]) Iterator() datastructures.Iterator[T] {
	return array.ModifyingIterator()
}

func (array *Array[T]) ModifyingIterator() sequences.Iterator[T] {
	return &iterator[T]{
		array: array,
		index: 0,
	}
}

func (array *Array[T]) Stream() *streams.Stream[T] {
	return streams.FromSlice(array.Slice())
}

func (array *Array[T]) Slice() []T {
	return *array
}
