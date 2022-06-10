package queue

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Queue[T any] *linkedlist.LinkedList[T]

func New[T any]() Queue[T] {
	return linkedlist.New[T]()
}

func (queue Queue[T]) getList() *linkedlist.LinkedList[T] {
	return queue
}

func (queue Queue[T]) Push(value T) {
	queue.getList().Append(value)
}

func (queue Queue[T]) Pop() T {
	return queue.getList().Remove(0)
}

func (queue Queue[T]) Peek() T {
	return queue.getList().Get(0)
}

func (queue Queue[T]) IsEmpty() bool {
	return queue.getList().Size() == 0
}
