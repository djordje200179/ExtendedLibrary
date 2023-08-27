package synclist

import "github.com/djordje200179/extendedlibrary/streams"

type collector[T any] struct {
	list *List[T]
}

func Collector[T any]() streams.Collector[T, *List[T]] {
	return &collector[T]{New[T]()}
}

func (c collector[T]) Supply(value T) {
	c.list.Prepend(value)
}

func (c collector[T]) Finish() *List[T] {
	return c.list
}
