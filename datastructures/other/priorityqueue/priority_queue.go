package priorityqueue

import (
	"github.com/djordje200179/GoExtendedLibrary/datastructures/sequences/linkedlist"
)

type Priority bool

const (
	SmallerFirst Priority = false
	BiggerFirst  Priority = true
)

type nodeData[T any] struct {
	value    T
	priority int
}

type PriorityQueue[T any] struct {
	list     linkedlist.LinkedList[nodeData[T]]
	priority Priority
}

func New[T any](priority Priority) PriorityQueue[T] {
	return PriorityQueue[T]{
		list:     linkedlist.New[nodeData[T]](),
		priority: priority,
	}
}

func (queue *PriorityQueue[T]) Push(value T) {
	for it := queue.list.Iterator(); it.IsValid(); it.Move() {
		// implement
	}
}

func (queue *PriorityQueue[T]) Pop() T {
	return queue.list.Remove(0).value
}

func (queue *PriorityQueue[T]) Peek() T {
	return queue.list.Get(0).value
}

func (queue *PriorityQueue[T]) IsEmpty() bool {
	return queue.list.Size() == 0
}
