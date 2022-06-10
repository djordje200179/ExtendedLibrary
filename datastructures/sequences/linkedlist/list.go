package linkedlist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func FromStream[T any](stream streams.Stream[T]) *LinkedList[T] {
	list := New[T]()

	stream.ForEach(func(value T) {
		list.Append(value)
	})

	return list
}

func (list *LinkedList[T]) getNode(index int) *node[T] {
	if index >= 0 {
		curr := list.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}

		return curr
	} else {
		index = -index - 1

		curr := list.tail
		for i := 0; i < index; i++ {
			curr = curr.prev
		}

		return curr
	}
}

func (list *LinkedList[T]) Size() int {
	return list.length
}

func (list *LinkedList[T]) Get(index int) T {
	return list.getNode(index).value
}

func (list *LinkedList[T]) Set(index int, value T) {
	list.getNode(index).value = value
}

func (list *LinkedList[T]) Append(values ...T) {
	for _, value := range values {
		newNode := &node[T]{value, list.tail, nil}

		if list.tail != nil {
			list.tail.next = newNode
			newNode.prev = list.tail
		} else {
			list.head = newNode
		}

		list.tail = newNode
	}

	list.length += len(values)
}

func (list *LinkedList[T]) Insert(index int, value T) {
	newNode := &node[T]{value, list.tail, nil}

	nextNode := list.getNode(index)
	prevNode := nextNode.prev

	if prevNode != nil {
		prevNode.next = newNode
	} else {
		list.head = newNode
	}

	if nextNode != nil {
		nextNode.prev = newNode
	} else {
		list.tail = newNode
	}
}

func (list *LinkedList[T]) Remove(index int) T {
	node := list.getNode(index)

	nextNode := node.next
	prevNode := node.prev

	if prevNode != nil {
		prevNode.next = nextNode
	} else {
		list.head = nextNode
	}

	if nextNode != nil {
		nextNode.prev = prevNode
	} else {
		list.tail = prevNode
	}

	list.length--

	return node.value
}

func (list *LinkedList[T]) Sort(comparator comparison.Comparator[T]) {
	// Implement
}

func (list *LinkedList[T]) Reverse() {
	// Implement
}

func (list *LinkedList[T]) Join(other sequences.Sequence[T]) {
	switch second := other.(type) {
	case *LinkedList[T]:
		list.tail.next = second.head
		second.head.prev = list.tail
		list.tail = second.tail

		list.length += second.length
	default:
		for it := other.Iterator(); it.IsValid(); it.Move() {
			list.Append(it.Get())
		}
	}
}

func (list *LinkedList[T]) Iterator() sequences.Iterator[T] {
	return &Iterator[T]{
		list:    list,
		current: list.head,
	}
}

func (list *LinkedList[T]) Stream() streams.Stream[T] {
	return sequences.CreateStream[T](list)
}
