package stack

import "github.com/djordje200179/extendedlibrary/streams"

type collector[T any] struct {
	stack Stack[T]
}

func Collector[T any](stack Stack[T]) streams.Collector[T, Stack[T]] {
	return collector[T]{stack}
}

func DefaultCollector[T any]() streams.Collector[T, Stack[T]] {
	return collector[T]{New[T]()}
}

func (collector collector[T]) Supply(value T) {
	collector.stack.Push(value)
}

func (collector collector[T]) Finish() Stack[T] {
	return collector.stack
}
