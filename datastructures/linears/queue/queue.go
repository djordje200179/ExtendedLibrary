package queue

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Queue[T any] struct {
	slice []T
}

func New[T any]() *Queue[T] {
	return NewWithCapacity[T](0)
}

func NewWithCapacity[T any](initialCapacity int) *Queue[T] {
	queue := new(Queue[T])
	queue.slice = make([]T, 0, initialCapacity)

	return queue
}

func (queue *Queue[T]) Push(value T) {
	queue.slice = append(queue.slice, value)
}

func (queue *Queue[T]) Pop() T {
	value := queue.Peek()
	queue.slice = queue.slice[1:]
	return value
}

func (queue *Queue[T]) Peek() T {
	return queue.slice[0]
}

func (queue *Queue[T]) Empty() bool {
	return len(queue.slice) == 0
}

func (queue *Queue[T]) ForEach(callback functions.ParamCallback[T]) {
	for !queue.Empty() {
		callback(queue.Pop())
	}
}
