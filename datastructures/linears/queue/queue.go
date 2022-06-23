package queue

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Queue[T any] array.Array[T]

func New[T any]() *Queue[T] {
	return NewWithCapacity[T](0)
}

func NewWithCapacity[T any](initialCapacity int) *Queue[T] {
	arr := array.NewWithCapacity[T](initialCapacity)
	return (*Queue[T])(arr)
}

func (queue *Queue[T]) array() *array.Array[T] {
	return (*array.Array[T])(queue)
}

func (queue *Queue[T]) Push(value T) {
	queue.array().Append(value)
}

func (queue *Queue[T]) Pop() T {
	defer queue.array().Remove(0)
	return queue.Peek()
}

func (queue *Queue[T]) Peek() T {
	return queue.array().Get(0)
}

func (queue *Queue[T]) Empty() bool {
	return queue.array().Size() == 0
}

func (queue *Queue[T]) ForEach(callback functions.ParamCallback[T]) {
	for !queue.Empty() {
		callback(queue.Pop())
	}
}
