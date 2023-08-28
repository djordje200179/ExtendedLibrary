package synclist

import "github.com/djordje200179/extendedlibrary/streams"

type collector[T any] struct {
	list *List[T]
}

// Collector creates a new stream Collector that collects elements into an empty List.
func Collector[T any]() streams.Collector[T, *List[T]] {
	return &collector[T]{New[T]()}
}

// Supply adds the specified value to the beginning of the List.
func (c collector[T]) Supply(value T) {
	c.list.Prepend(value)
}

// Finish returns the List that was built.
func (c collector[T]) Finish() *List[T] {
	return c.list
}
