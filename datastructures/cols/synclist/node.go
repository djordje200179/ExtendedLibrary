package synclist

// Node is a node in a singly linked list.
type Node[T any] struct {
	Value T // The value stored in the node.

	next *Node[T]
}

// Next returns the next node in the list.
func (node Node[T]) Next() *Node[T] {
	return node.next
}
