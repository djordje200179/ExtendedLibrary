package linkedlist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

type List[T any] struct {
	head, tail *Node[T]
	size       int
}

func New[T any]() *List[T] {
	return new(List[T])
}

func Collector[T any]() streams.Collector[T, *List[T]] {
	return collections.Collector[T, *List[T]]{New[T]()}
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) GetNode(index int) *Node[T] {
	if index >= list.size || index < -list.size {
		collections.PanicOnIndexOutOfBounds(index, list.size)
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

func (list *List[T]) Prepend(value T) {
	node := &Node[T]{
		Value: value,
	}

	node.next = list.head
	if list.head == nil {
		list.head.prev = node
	} else {
		list.tail = node
	}
	list.head = node

	list.size++
}

func (list *List[T]) Append(value T) {
	if list.size == 0 {
		node := &Node[T]{
			Value: value,
		}

		list.head = node
		list.tail = node
		list.size++
	} else {
		list.InsertAfterNode(list.tail, value)
	}
}

func (list *List[T]) Insert(index int, value T) {
	prevNode := list.GetNode(index)
	list.InsertBeforeNode(prevNode, value)
}

func (list *List[T]) InsertBeforeNode(nextNode *Node[T], value T) {
	newNode := &Node[T]{value, nextNode.prev, nextNode}

	if nextNode.prev != nil {
		nextNode.prev.next = newNode
	} else {
		list.head = newNode
	}

	nextNode.prev = newNode
	list.size++
}

func (list *List[T]) InsertAfterNode(prevNode *Node[T], value T) {
	newNode := &Node[T]{value, prevNode, prevNode.next}

	if prevNode.next != nil {
		prevNode.next.prev = newNode
	} else {
		list.tail = newNode
	}

	prevNode.next = newNode
	list.size++
}

func (list *List[T]) Remove(index int) {
	node := list.GetNode(index)
	list.RemoveNode(node)
}

func (list *List[T]) RemoveNode(node *Node[T]) {
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

func (list *List[T]) SwapNodes(node1, node2 *Node[T]) {
	if node1 == node2 {
		return
	}

	if (node1.prev == nil && list.head != node1) || (node1.next == nil && list.tail != node1) {
		panic("node1 is not part of the list")
	}

	if (node2.prev == nil && list.head != node2) || (node2.next == nil && list.tail != node2) {
		panic("node2 is not part of the list")
	}

	node1.prev, node2.prev = node2.prev, node1.prev
	node1.next, node2.next = node2.next, node1.next

	if node1.prev != nil {
		node1.prev.next = node1
	} else {
		list.head = node1
	}

	if node1.next != nil {
		node1.next.prev = node1
	} else {
		list.tail = node1
	}

	if node2.prev != nil {
		node2.prev.next = node2
	} else {
		list.head = node2
	}

	if node2.next != nil {
		node2.next.prev = node2
	} else {
		list.tail = node2
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
	return list.CollectionIterator()
}

func (list *List[T]) CollectionIterator() collections.Iterator[T] {
	return &Iterator[T]{
		list:  list,
		curr:  list.head,
		index: 0,
	}
}

func (list *List[T]) Stream() streams.Stream[T] {
	return iterable.IteratorStream(list.Iterator())
}

func (list *List[T]) RefsStream() streams.Stream[*T] {
	return collections.RefsStream(list.CollectionIterator())
}

func (list *List[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	for curr, i := list.head, 0; curr != nil; curr, i = curr.next, i+1 {
		if predicate(curr.Value) {
			return i, true
		}
	}

	return -1, false
}

func (list *List[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	node, ok := list.FindNode(predicate)
	if !ok {
		return nil, false
	}

	return &node.Value, true
}

func (list *List[T]) FindNode(predicate predication.Predicate[T]) (*Node[T], bool) {
	for curr := list.head; curr != nil; curr = curr.next {
		if predicate(curr.Value) {
			return curr, true
		}
	}

	return nil, false
}

func (list *List[T]) Head() *Node[T] {
	return list.head
}

func (list *List[T]) Tail() *Node[T] {
	return list.tail
}
