package priorityqueue

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T any] struct {
	queue PriorityQueue[T]
}

func Collector[T any](priority Priority) streams.Collector[misc.Pair[T, int], PriorityQueue[T]] {
	return collector[T]{New[T](priority)}
}

func (collector collector[T]) Supply(value misc.Pair[T, int]) {
	collector.queue.Push(value.First, value.Second)
}
func (collector collector[T]) Finish() PriorityQueue[T] { return collector.queue }
