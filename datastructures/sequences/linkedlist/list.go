package linkedlist

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func Collector[T any]() streams.Collector[T, sequences.Sequence[T]] {
	return sequences.Collector[T](New[T]())
}

func (list *LinkedList[T]) getNode(index int) *node[T] {
	if index >= list.size || index < -list.size {
		panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", index, list.size))
	}

	if index < 0 {
		index = -index - 1
	}

	var curr *node[T]
	if index >= 0 {
		curr = list.head
	} else {
		curr = list.tail
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

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) Get(index int) T {
	return list.getNode(index).value
}

func (list *LinkedList[T]) Set(index int, value T) {
	list.getNode(index).value = value
}

func (list *LinkedList[T]) Append(value T) {
	newNode := &node[T]{value, list.tail, nil}

	if list.tail != nil {
		list.tail.next = newNode
		newNode.prev = list.tail
	} else {
		list.head = newNode
	}
	list.tail = newNode

	list.size++
}

func (list *LinkedList[T]) AppendMany(values ...T) {
	for _, value := range values {
		list.Append(value)
	}
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

func (list *LinkedList[T]) Remove(index int) {
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

	list.size--
}

func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *LinkedList[T]) Sort(comparator functions.Comparator[T]) {
	for front := list.head; front.next != nil; front = front.next {
		for back := front.next; back != nil; back = back.next {
			if comparator(front.value, back.value) != comparison.FirstSmaller {
				front.value, back.value = back.value, front.value
			}
		}
	}
}

func (list *LinkedList[T]) Join(other sequences.Sequence[T]) {
	switch second := other.(type) {
	case *LinkedList[T]:
		list.tail.next = second.head
		second.head.prev = list.tail
		list.tail = second.tail

		list.size += second.size
	default:
		for it := other.Iterator(); it.Valid(); it.Move() {
			list.Append(it.Get())
		}
	}
}

func (list *LinkedList[T]) Clone() sequences.Sequence[T] {
	cloned := New[T]()
	for it := list.Iterator(); it.Valid(); it.Move() {
		cloned.Append(it.Get())
	}

	return cloned
}

func (list *LinkedList[T]) Iterator() datastructures.Iterator[T] {
	return list.ModifyingIterator()
}

func (list *LinkedList[T]) ModifyingIterator() sequences.Iterator[T] {
	return &Iterator[T]{
		list:    list,
		current: list.head,
	}
}

func (list *LinkedList[T]) Stream() streams.Stream[T] {
	return streams.FromIterable[T](list)
}
