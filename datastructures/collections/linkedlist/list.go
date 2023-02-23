package linkedlist

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type List[T any] struct {
	head, tail *Node[T]
	size       int
}

func New[T any]() *List[T] {
	return new(List[T])
}

func Collector[T any]() streams.Collector[T, collections.Collection[T]] {
	return collections.Collector[T]{New[T]()}
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) GetNode(index int) *Node[T] {
	if index >= list.size || index < -list.size {
		panic(fmt.Sprintf("Index out of bounds: %d", index))
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

func (list *List[T]) GetRef(index int) *T {
	return &list.GetNode(index).Value
}

func (list *List[T]) Get(index int) T {
	return *list.GetRef(index)
}

func (list *List[T]) Set(index int, value T) {
	*list.GetRef(index) = value
}

func (list *List[T]) Append(values ...T) {
	for _, value := range values {
		if list.size == 0 {
			node := &Node[T]{
				Value: value,
			}

			list.head = node
			list.tail = node
			list.size++
		} else {
			list.insertAfter(list.tail, value)
		}
	}
}

func (list *List[T]) Insert(index int, values ...T) {
	for _, value := range values {
		list.insertBefore(list.GetNode(index), value)
	}
}

func (list *List[T]) Remove(index int) {
	list.removeNode(list.GetNode(index))
}

func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *List[T]) Reverse() {
	for curr := list.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
	}

	list.head, list.tail = list.tail, list.head
}

func (list *List[T]) Swap(index1, index2 int) {
	node1 := list.GetNode(index1)
	node2 := list.GetNode(index2)

	node1.Value, node2.Value = node2.Value, node1.Value
}

func (list *List[T]) Sort(comparator functions.Comparator[T]) {
	for front := list.head; front.next != nil; front = front.next {
		for back := front.next; back != nil; back = back.next {
			if comparator(front.Value, back.Value) != comparison.FirstSmaller {
				front.Value, back.Value = back.Value, front.Value
			}
		}
	}
}

func (list *List[T]) Join(other collections.Collection[T]) {
	switch second := other.(type) {
	case *List[T]:
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

func (list *List[T]) Clone() collections.Collection[T] {
	cloned := New[T]()
	for curr := list.head; curr != nil; curr = curr.next {
		cloned.Append(curr.Value)
	}

	return cloned
}

func (list *List[T]) Iterator() iterable.Iterator[T] {
	return list.ModifyingIterator()
}

func (list *List[T]) ModifyingIterator() collections.Iterator[T] {
	return &Iterator[T]{
		list:  list,
		curr:  list.head,
		index: 0,
	}
}

func (list *List[T]) Stream() streams.Stream[T] {
	supplier := iterable.IteratorSupplier[T]{list.Iterator()}
	return streams.Stream[T]{supplier}
}

func (list *List[T]) RefStream() streams.Stream[*T] {
	supplier := collections.RefsSupplier[T]{list.ModifyingIterator()}
	return streams.Stream[*T]{supplier}
}

func (list *List[T]) Head() *Node[T] {
	return list.head
}

func (list *List[T]) Tail() *Node[T] {
	return list.tail
}
