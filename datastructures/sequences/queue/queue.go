package queue

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/collections/array"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Queue[T any] struct {
	sequence collections.Collection[T]
}

func NewFrom[T any](sequence collections.Collection[T]) Queue[T] { return Queue[T]{sequence} }
func New[T any]() Queue[T]                                       { return Queue[T]{array.New[T]()} }

func (queue Queue[T]) Empty() bool { return queue.sequence.Size() == 0 }

func (queue Queue[T]) Push(value T) { queue.sequence.Append(value) }
func (queue Queue[T]) Peek() T      { return queue.sequence.Get(0) }
func (queue Queue[T]) Pop() T {
	defer queue.sequence.Remove(0)
	return queue.Peek()
}

func (queue Queue[T]) ForEach(callback functions.ParamCallback[T]) {
	for !queue.Empty() {
		callback(queue.Pop())
	}
}
