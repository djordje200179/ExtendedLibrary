package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Array[T any] []T

func New[T any](initialCapacity int) *Array[T] {
	arr := make([]T, 0, initialCapacity)
	return (*Array[T])(&arr)
}

func (array *Array[T]) Size() int {
	return len(*array)
}

func (array *Array[T]) Get(i int) T {
	return (*array)[i]
}

func (array *Array[T]) Set(i int, value T) {
	(*array)[i] = value
}

func (array *Array[T]) Append(values ...T) {
	*array = append(*array, values...)
}

func (array *Array[T]) Insert(index int, value T) {
	*array = append((*array)[:index+1], (*array)[index:]...)
	(*array)[index] = value
}

func (array *Array[T]) Remove(index int) {
	*array = append((*array)[:index], (*array)[index+1:]...)
}

func (array *Array[T]) Empty() {
	*array = nil
}

func (array *Array[T]) Sort(comparator functions.Comparator[T]) {
	// Implement
}

func (array *Array[T]) Join(other sequences.Sequence[T]) {
	// Implement
}

func (array *Array[T]) Iterator() sequences.Iterator[T] {
	return &iterator[T]{
		array: array,
		index: 0,
	}
}

func (array *Array[T]) Stream() streams.Stream[T] {
	return streams.FromSlice(*array)
}
