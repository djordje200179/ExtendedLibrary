package linkedlist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type LinkedList[T any] struct {
	head, tail *node[T]
	size       int
}

func New[T any]() *LinkedList[T] {
	return new(LinkedList[T])
}

func Collector[T any]() streams.Collector[T, collections.Collection[T]] {
	return collections.Collector[T]{
		Collection: New[T](),
	}
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) GetRef(index int) *T {
	return &list.getNode(index).value
}

func (list *LinkedList[T]) Get(index int) T {
	return *list.GetRef(index)
}

func (list *LinkedList[T]) Set(index int, value T) {
	*list.GetRef(index) = value
}

func (list *LinkedList[T]) Append(values ...T) {
	for _, value := range values {
		if list.size == 0 {
			node := &node[T]{value: value}
			list.head = node
			list.tail = node
			list.size++
		} else {
			list.insertAfterNode(list.tail, value)
		}
	}
}

func (list *LinkedList[T]) Insert(index int, values ...T) {
	for _, value := range values {
		list.insertBeforeNode(list.getNode(index), value)
	}
}

func (list *LinkedList[T]) Remove(index int) {
	list.removeNode(list.getNode(index))
}

func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *LinkedList[T]) Reverse() {
	for curr := list.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
	}

	list.head, list.tail = list.tail, list.head
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

func (list *LinkedList[T]) Join(other collections.Collection[T]) {
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

	other.Clear()
}

func (list *LinkedList[T]) Clone() collections.Collection[T] {
	cloned := New[T]()
	for curr := list.head; curr != nil; curr = curr.next {
		cloned.Append(curr.value)
	}

	return cloned
}

func (list *LinkedList[T]) Iterator() iterable.Iterator[T] {
	return list.ModifyingIterator()
}

func (list *LinkedList[T]) ModifyingIterator() collections.Iterator[T] {
	return &Iterator[T]{
		list:  list,
		curr:  list.head,
		index: 0,
	}
}

func (list *LinkedList[T]) Stream() streams.Stream[T] {
	return collections.ValuesStream[T](list)
}

func (list *LinkedList[T]) RefStream() streams.Stream[*T] {
	return collections.RefsStream[T](list)
}
