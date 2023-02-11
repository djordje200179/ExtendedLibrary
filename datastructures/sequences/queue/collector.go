package queue

import (
	"github.com/djordje200179/extendedlibrary/streams"
)

type collector[T any] struct {
	queue Queue[T]
}

func Collector[T any](queue Queue[T]) streams.Collector[T, Queue[T]] { return collector[T]{queue} }
func DefaultCollector[T any]() streams.Collector[T, Queue[T]]        { return collector[T]{New[T]()} }

func (collector collector[T]) Supply(value T)   { collector.queue.Push(value) }
func (collector collector[T]) Finish() Queue[T] { return collector.queue }
