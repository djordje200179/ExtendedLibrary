package linkedlist

import "fmt"

type Node[T any] struct {
	Value      T
	prev, next *Node[T]
}

func (node *Node[T]) Prev() *Node[T] { return node.prev }
func (node *Node[T]) Next() *Node[T] { return node.next }

func (list *LinkedList[T]) Head() *Node[T] { return list.head }
func (list *LinkedList[T]) Tail() *Node[T] { return list.tail }

func (list *LinkedList[T]) GetNode(index int) *Node[T] {
	if index >= list.size || index < -list.size {
		//TODO: Improve panic type
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, list.size))
	}

	var curr *Node[T]
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

func (list *LinkedList[T]) insertBeforeNode(node *Node[T], value T) {
	newNode := &Node[T]{value, node.prev, node}

	if node.prev != nil {
		node.prev.next = newNode
	} else {
		list.head = newNode
	}

	node.prev = newNode
	list.size++
}

func (list *LinkedList[T]) insertAfterNode(node *Node[T], value T) {
	newNode := &Node[T]{value, node, node.next}

	if node.next != nil {
		node.next.prev = newNode
	} else {
		list.tail = newNode
	}

	node.next = newNode
	list.size++
}

func (list *LinkedList[T]) removeNode(node *Node[T]) {
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
