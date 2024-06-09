package synclist

// Node is an element of a List.
// It should not be created directly.
type Node[T any] struct {
	Value T // Stored value

	next *Node[T]
}

// Next returns the next Node.
func (node Node[T]) Next() *Node[T] {
	return node.next
}
