package linklist

// Node is an element of a List.
// It should not be created directly.
type Node[T any] struct {
	Value T // Stored value

	list *List[T]

	prev, next *Node[T]
}

// Prev returns the previous Node.
func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}

// Next returns the next Node.
func (node *Node[T]) Next() *Node[T] {
	return node.next
}

// InsertBefore inserts the specified element
// before this node.
func (node *Node[T]) InsertBefore(value T) {
	newNode := &Node[T]{value, node.list, node.prev, node}

	if node.prev != nil {
		node.prev.next = newNode
	} else {
		node.list.head = newNode
	}

	node.prev = newNode
	node.list.size++
}

// InsertAfter inserts the specified element
// after this node.
func (node *Node[T]) InsertAfter(value T) {
	newNode := &Node[T]{value, node.list, node, node.next}

	if node.next != nil {
		node.next.prev = newNode
	} else {
		node.list.tail = newNode
	}

	node.next = newNode
	node.list.size++
}
