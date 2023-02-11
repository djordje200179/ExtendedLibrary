package linkedlist

import "fmt"

type node[T any] struct {
	value T

	prev, next *node[T]
}

func (list *LinkedList[T]) getNode(index int) *node[T] {
	if index >= list.size || index < -list.size {
		//TODO: Improve panic type
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, list.size))
	}

	var curr *node[T]
	if index >= 0 {
		curr = list.head
	} else {
		curr = list.tail
	}

	if index < 0 {
		index = -index - 1
	}

	for i := 0; i < index; i++ {
		if index >= 0 {
			curr = curr.next
		} else {
			curr = curr.prev
		}
	}

	return curr
}

func (list *LinkedList[T]) insertBeforeNode(nextNode *node[T], value T) {
	newNode := &node[T]{value, nextNode.prev, nextNode}

	if nextNode.prev != nil {
		nextNode.prev.next = newNode
	} else {
		list.head = newNode
	}

	nextNode.prev = newNode
	list.size++
}

func (list *LinkedList[T]) insertAfterNode(prevNode *node[T], value T) {
	newNode := &node[T]{value, prevNode, prevNode.next}

	if prevNode.next != nil {
		prevNode.next.prev = newNode
	} else {
		list.tail = newNode
	}

	prevNode.next = newNode
	list.size++
}

func (list *LinkedList[T]) removeNode(node *node[T]) {
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
