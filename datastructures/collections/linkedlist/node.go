package linkedlist

type Node[T any] struct {
	Value T

	list *List[T]

	prev, next *Node[T]
}

func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

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
