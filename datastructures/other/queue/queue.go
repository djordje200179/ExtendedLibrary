package queue

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Queue[T any] struct {
	seq sequences.Sequence[T]
}

func New[T any]() Queue[T] {
	return Queue[T]{linkedlist.New[T]()}
}

func (queue Queue[T]) Push(value T) {
	queue.seq.Append(value)
}

func (queue Queue[T]) Pop() T {
	defer queue.seq.Remove(0)
	return queue.Peek()
}

func (queue Queue[T]) Peek() T {
	return queue.seq.Get(0)
}

func (queue Queue[T]) Empty() bool {
	return queue.seq.Size() == 0
}
