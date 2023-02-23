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

func (list *List[T]) insertBefore(nextNode *Node[T], value T) {
	newNode := &Node[T]{value, nextNode.prev, nextNode}

	if nextNode.prev != nil {
		nextNode.prev.next = newNode
	} else {
		list.head = newNode
	}

	nextNode.prev = newNode
	list.size++
}

func (list *List[T]) insertAfter(prevNode *Node[T], value T) {
	newNode := &Node[T]{value, prevNode, prevNode.next}

	if prevNode.next != nil {
		prevNode.next.prev = newNode
	} else {
		list.tail = newNode
	}

	prevNode.next = newNode
	list.size++
}

func (list *List[T]) removeNode(node *Node[T]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}

	list.size--
}
