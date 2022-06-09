package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Array[T any] []T

func New[T any](capacity int) Array[T] {
	return make([]T, 0, capacity)
}

func FromStream[T any](stream streams.Stream[T]) Array[T] {
	array := New[T](0)

	stream.ForEach(func(value T) {
		array.Append(value)
	})

	return array
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

func (array *Array[T]) Remove(index int) T {
	value := (*array)[index]

	*array = append((*array)[:index], (*array)[index+1:]...)

	return value
}

func (array *Array[T]) Sort(comparator comparison.Comparator[T]) {
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
	return streams.FromValues(([]T)(*array)...)
}
