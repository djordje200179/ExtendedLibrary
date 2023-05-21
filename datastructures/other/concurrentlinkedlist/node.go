package concurrentlinkedlist

type Node[T any] struct {
	Value T

	next *Node[T]
}

func (node Node[T]) Next() *Node[T] {
	return node.next
}
