package linkedlist

import stream "github.com/djordje200179/extendedlibrary/streams/collector"

type collector[T any] struct {
	list *LinkedList[T]
}

func Collector[T any]() stream.Collector[T, *LinkedList[T]] {
	return collector[T]{
		list: New[T](),
	}
}

func (c collector[T]) Supply(value T) {
	c.list.Append(value)
}

func (c collector[T]) Finish() *LinkedList[T] {
	return c.list
}
