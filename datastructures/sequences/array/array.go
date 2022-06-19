package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
	"sort"
)

type Array[T any] struct {
	slice []T
}

func New[T any](initialCapacity int) *Array[T] {
	return &Array[T]{
		slice: make([]T, 0, initialCapacity),
	}
}

func Collector[T any]() streams.Collector[T, sequences.Sequence[T]] {
	return sequences.Collector[T](New[T](0))
}

func (array *Array[T]) Size() int {
	return len(array.slice)
}

func (array *Array[T]) Get(i int) T {
	return array.slice[i]
}

func (array *Array[T]) Set(i int, value T) {
	array.slice[i] = value
}

func (array *Array[T]) Append(value T) {
	array.slice = append(array.slice, value)
}

func (array *Array[T]) AppendMany(values ...T) {
	array.slice = append(array.slice, values...)
}

func (array *Array[T]) Insert(index int, value T) {
	array.slice = append(array.slice[:index+1], array.slice[index:]...)
	array.slice[index] = value
}

func (array *Array[T]) Remove(index int) {
	array.slice = append(array.slice[:index], array.slice[index+1:]...)
}

func (array *Array[T]) Clear() {
	array.slice = nil
}

func (array *Array[T]) Sort(comparator functions.Comparator[T]) {
	sort.Slice(array.slice, func(i, j int) bool {
		return comparator(array.slice[i], array.slice[j]) == comparison.FirstSmaller
	})
}

func (array *Array[T]) Join(other sequences.Sequence[T]) {
	switch second := other.(type) {
	case *Array[T]:
		array.AppendMany(second.slice...)
	default:
		for it := other.Iterator(); it.Valid(); it.Move() {
			array.Append(it.Get())
		}
	}
}

func (array *Array[T]) Clone() sequences.Sequence[T] {
	cloned := Array[T]{
		slice: make([]T, len(array.slice)),
	}
	copy(cloned.slice, array.slice)

	return &cloned
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
	return streams.FromSlice(array.slice)
}
