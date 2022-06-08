package queue

import (
	"github.com/djordje200179/GoExtendedLibrary/datastructures/sequences/linkedlist"
)

type Queue[T any] struct {
	list linkedlist.LinkedList[T]
}

func New[T any]() Queue[T] {
	return Queue[T]{linkedlist.New[T]()}
}

func (queue *Queue[T]) Push(value T) {
	queue.list.Append(value)
}

func (queue *Queue[T]) Pop() T {
	return queue.list.Remove(0)
}

func (queue *Queue[T]) Peek() T {
	return queue.list.Get(0)
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.Size() == 0
}
