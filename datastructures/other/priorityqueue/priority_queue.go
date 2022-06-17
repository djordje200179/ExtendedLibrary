package priorityqueue

import (
	"container/heap"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Priority bool

const (
	SmallerFirst Priority = false
	BiggerFirst  Priority = true
)

type PriorityQueue[T any] struct {
	heapSlice[T]
}

func New[T any](priority Priority) *PriorityQueue[T] {
	return &PriorityQueue[T]{heapSlice[T]{
		slice:    nil,
		priority: priority,
	}}
}

func (queue *PriorityQueue[T]) Push(value T, priority int) {
	item := item[T]{
		value:    value,
		priority: priority,
	}

	heap.Push(&queue.heapSlice, item)
}

func (queue *PriorityQueue[T]) Pop() T {
	item := heap.Pop(&queue.heapSlice).(item[T])
	return item.value
}

func (queue *PriorityQueue[T]) Peek() T {
	return queue.heapSlice.slice[0].value
}

func (queue *PriorityQueue[T]) Empty() bool {
	return queue.heapSlice.Len() == 0
}

func (queue *PriorityQueue[T]) ForEach(callback functions.ParamCallback[T]) {
	for !queue.Empty() {
		callback(queue.Pop())
	}
}
