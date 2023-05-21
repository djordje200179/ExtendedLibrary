package concurrentlinkedlist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/streams"
	"sync/atomic"
)

type List[T any] struct {
	head atomic.Pointer[Node[T]]
}

func New[T any]() *List[T] {
	return new(List[T])
}

func (list *List[T]) Prepend(value T) {
	node := &Node[T]{Value: value}

	node.next = list.head.Swap(node)
}

func (list *List[T]) Iterator() iterable.Iterator[T] {
	return &Iterator[T]{curr: list.head.Load()}
}

func (list *List[T]) Stream() streams.Stream[T] {
	return iterable.IteratorStream(list.Iterator())
}

func (list *List[T]) Head() *Node[T] {
	return list.head.Load()
}
