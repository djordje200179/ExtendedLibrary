package linkedlist

type Node[T any] struct {
	Value T

	prev, next *Node[T]
}

func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}
