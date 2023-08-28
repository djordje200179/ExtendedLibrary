package linklist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

// List is a doubly linked list implementation.
// The zero value is ready to use. Do not copy a non-zero List.
type List[T any] struct {
	head, tail *Node[T]
	size       int
}

// New creates an empty list.
func New[T any]() *List[T] {
	return new(List[T])
}

// NewFromIterable creates a list from the specified iter.Iterable.
func NewFromIterable[T any](iterable iter.Iterable[T]) *List[T] {
	list := New[T]()

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		list.Append(it.Get())
	}

	return list
}

// Size returns the number of elements in the list.
func (list *List[T]) Size() int {
	return list.size
}

// GetNode returns the node at the specified index.
// Negative indices are interpreted as relative to the end of the List.
// Panics if the index is out of bounds.
func (list *List[T]) GetNode(index int) *Node[T] {
	if index >= list.size || index < -list.size {
		panic(cols.ErrIndexOutOfBounds{Index: index, Length: list.size})
	}

	if index == 0 {
		return list.head
	}

	var curr *Node[T]
	if index > 0 {
		curr = list.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}
	} else {
		curr = list.tail
		for i := -1; i > index; i-- {
			curr = curr.prev
		}
	}
	return curr
}

// GetRef returns a reference to the element at the specified index.
// Negative indices are interpreted as relative to the end of the List.
// Panics if the index is out of bounds.
func (list *List[T]) GetRef(index int) *T {
	node := list.GetNode(index)
	return &node.Value
}

// Get returns the element at the specified index.
// Negative indices are interpreted as relative to the end of the List.
// Panics if the index is out of bounds.
func (list *List[T]) Get(index int) T {
	node := list.GetNode(index)
	return node.Value
}

// Set sets the element at the specified index.
// Negative indices are interpreted as relative to the end of the List.
// Panics if the index is out of bounds.
func (list *List[T]) Set(index int, value T) {
	node := list.GetNode(index)
	node.Value = value
}

// Prepend inserts the specified element at the beginning of the List.
func (list *List[T]) Prepend(value T) {
	node := &Node[T]{
		Value: value,
	}

	node.next = list.head
	if list.head != nil {
		list.head.prev = node
	} else {
		list.tail = node
	}
	list.head = node

	list.size++
}

// Append inserts the specified element at the end of the List.
func (list *List[T]) Append(value T) {
	if list.size == 0 {
		node := &Node[T]{
			Value: value,
		}

		list.head = node
		list.tail = node
		list.size++
	} else {
		list.tail.InsertAfter(value)
	}
}

// Insert inserts the specified element at the specified index.
// Negative indices are interpreted as relative to the end of the List.
// Panics if the index is out of bounds.
func (list *List[T]) Insert(index int, value T) {
	list.GetNode(index).InsertBefore(value)
}

// Remove removes the element at the specified index.
// Negative indices are interpreted as relative to the end of the List.
// Panics if the index is out of bounds.
func (list *List[T]) Remove(index int) {
	node := list.GetNode(index)
	list.RemoveNode(node)
}

// RemoveNode removes the specified node from the List.
// Panics if the node is not part of the List.
func (list *List[T]) RemoveNode(node *Node[T]) {
	if node == nil || node.list != list {
		panic("invalid node")
	}

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

// Clear removes all elements from the List.
func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Reverse reverses the order of the elements in the List.
func (list *List[T]) Reverse() {
	for curr := list.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
	}

	list.head, list.tail = list.tail, list.head
}

// Join appends all elements from the other collection to the List.
// The other collection is cleared.
func (list *List[T]) Join(other cols.Collection[T]) {
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

// Clone returns a shallow copy of the List.
func (list *List[T]) Clone() cols.Collection[T] {
	cloned := New[T]()
	for curr := list.head; curr != nil; curr = curr.next {
		cloned.Append(curr.Value)
	}

	return cloned
}

// Iterator returns an iterator over the elements in the List.
func (list *List[T]) Iterator() iter.Iterator[T] {
	return list.CollectionIterator()
}

// CollectionIterator returns an iterator over the elements in the List.
func (list *List[T]) CollectionIterator() cols.Iterator[T] {
	return &Iterator[T]{
		list:  list,
		curr:  list.head,
		index: 0,
	}
}

// Stream returns a stream of the elements in the List.
func (list *List[T]) Stream() streams.Stream[T] {
	return iter.IteratorStream(list.Iterator())
}

// RefsStream returns a stream of references to the elements in the List.
func (list *List[T]) RefsStream() streams.Stream[*T] {
	return cols.RefsStream(list.CollectionIterator())
}

// FindIndex returns the index of the first element that matches the specified predicate.
func (list *List[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	for curr, i := list.head, 0; curr != nil; curr, i = curr.next, i+1 {
		if predicate(curr.Value) {
			return i, true
		}
	}

	return -1, false
}

// FindRef returns a reference to the first element that matches the specified predicate.
func (list *List[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	node, ok := list.FindNode(predicate)
	if !ok {
		return nil, false
	}

	return &node.Value, true
}

// FindNode returns the first node that matches the specified predicate.
func (list *List[T]) FindNode(predicate predication.Predicate[T]) (*Node[T], bool) {
	for curr := list.head; curr != nil; curr = curr.next {
		if predicate(curr.Value) {
			return curr, true
		}
	}

	return nil, false
}

// Head returns the first node in the List.
func (list *List[T]) Head() *Node[T] {
	return list.head
}

// Tail returns the last node in the List.
func (list *List[T]) Tail() *Node[T] {
	return list.tail
}
