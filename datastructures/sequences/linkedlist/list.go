package linkedlist

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type LinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

func New[T any]() *LinkedList[T] { return new(LinkedList[T]) }

func NewWithSize[T any](initialSize int) *LinkedList[T] {
	list := New[T]()

	for i := 0; i < initialSize; i++ {
		var zeroValue T
		list.Append(zeroValue)
	}

	return list
}

func Collector[T any]() streams.Collector[T, sequences.Sequence[T]] {
	return sequences.Collector[T](New[T]())
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) GetRef(index int) *T    { return &list.GetNode(index).Value }
func (list *LinkedList[T]) Get(index int) T        { return *list.GetRef(index) }
func (list *LinkedList[T]) Set(index int, value T) { *list.GetRef(index) = value }

func (list *LinkedList[T]) Append(value T) {
	if list.size == 0 {
		node := &Node[T]{Value: value}
		list.head = node
		list.tail = node
		list.size++
	} else {
		list.insertAfterNode(list.tail, value)
	}
}

func (list *LinkedList[T]) AppendMany(values ...T) {
	for _, value := range values {
		list.Append(value)
	}
}

func (list *LinkedList[T]) Insert(index int, value T) {
	list.insertBeforeNode(list.GetNode(index), value)
}

func (list *LinkedList[T]) Remove(index int) {
	list.removeNode(list.GetNode(index))
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
			if comparator(front.Value, back.Value) != comparison.FirstSmaller {
				front.Value, back.Value = back.Value, front.Value
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
	for curr := list.head; curr != nil; curr = curr.next {
		cloned.Append(curr.Value)
	}

	return cloned
}

func (list *LinkedList[T]) Iterator() datastructures.Iterator[T] {
	return list.ModifyingIterator()
}

func (list *LinkedList[T]) ModifyingIterator() sequences.Iterator[T] {
	return &Iterator[T]{
		list: list,
		curr: list.head,
	}
}

func (list *LinkedList[T]) Stream() *streams.Stream[T] {
	return streams.FromIterable[T](list)
}

func (list *LinkedList[T]) RefStream() *streams.Stream[*T] {
	iterator := list.ModifyingIterator()
	return streams.Supply(func() *T {
		defer iterator.Move()
		return iterator.GetRef()
	})
}
