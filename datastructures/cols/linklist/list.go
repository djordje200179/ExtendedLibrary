package linklist

import (
	"github.com/djordje200179/extendedlibrary/datastructures/cols"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

// List is a cols.Collection that
// stores elements in a doubly linked list.
//
// The zero value is ready to use.
// Do not copy a non-zero List.
type List[T any] struct {
	head, tail *Node[T]
	size       int
}

// New creates an empty List.
func New[T any]() *List[T] {
	return new(List[T])
}

// NewFromIterable creates a new List from the specified iter.Iterable.
func NewFromIterable[T any](iterable iter.Iterable[T]) *List[T] {
	list := New[T]()

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		val := it.Get()
		list.Append(val)
	}

	return list
}

// Size returns the number of elements.
func (list *List[T]) Size() int {
	return list.size
}

// GetNode returns the Node at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (list *List[T]) GetNode(index int) *Node[T] {
	if index >= list.size || index < -list.size {
		panic(cols.IndexOutOfBoundsError{Index: index, Length: list.size})
	}

	if index == 0 {
		return list.head
	}

	var curr *Node[T]
	if index > 0 {
		curr = list.head
		for range index {
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
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (list *List[T]) GetRef(index int) *T { return &list.GetNode(index).Value }

// Get returns the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (list *List[T]) Get(index int) T { return list.GetNode(index).Value }

// Set sets the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (list *List[T]) Set(index int, value T) { list.GetNode(index).Value = value }

// Prepend inserts the specified element at the beginning.
func (list *List[T]) Prepend(value T) {
	if list.size == 0 {
		node := &Node[T]{
			Value: value,
			list:  list,
		}

		list.head = node
		list.tail = node
		list.size++
	} else {
		list.head.InsertBefore(value)
	}
}

// Append appends the specified element to the end.
func (list *List[T]) Append(value T) {
	if list.size == 0 {
		node := &Node[T]{
			Value: value,
			list:  list,
		}

		list.head = node
		list.tail = node
		list.size++
	} else {
		list.tail.InsertAfter(value)
	}
}

// Insert inserts the specified element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (list *List[T]) Insert(index int, value T) {
	list.GetNode(index).InsertBefore(value)
}

// Remove removes the element at the specified index.
//
// Negative indices are interpreted as relative to the end.
// Panic occurs if the index is out of bounds.
func (list *List[T]) Remove(index int) {
	node := list.GetNode(index)
	list.RemoveNode(node)
}

// RemoveNode removes the element associated with the specified node.
//
// Panic occurs if the node is not associated with the List.
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

// Clear removes all elements.
func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Reverse reverses the order of the elements.
func (list *List[T]) Reverse() {
	for curr := list.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
	}

	list.head, list.tail = list.tail, list.head
}

// Join moves all elements from the other cols.Collection
// to the end. The other cols.Collection becomes empty.
//
// If the other collection is a List, only pointers to the nodes are moved.
// Unfortunately, it is also needed to update the reference
// to the list in the moved nodes.
// If the other collection is not a List, the elements are moved one by one.
func (list *List[T]) Join(other cols.Collection[T]) {
	switch second := other.(type) {
	case *List[T]:
		for curr := second.head; curr != nil; curr = curr.next {
			curr.list = list
		}

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

// Clone returns a copy of the List.
func (list *List[T]) Clone() cols.Collection[T] {
	cloned := New[T]()
	for curr := list.head; curr != nil; curr = curr.next {
		cloned.Append(curr.Value)
	}

	return cloned
}

// Iterator returns a read-only iter.Iterator over the elements.
//
// Iteration starts from the first element.
func (list *List[T]) Iterator() iter.Iterator[T] {
	return list.CollectionIterator()
}

// CollectionIterator returns an Iterator over the elements.
// It can be used to modify the elements while iterating.
//
// Iteration starts from the first element.
func (list *List[T]) CollectionIterator() cols.Iterator[T] {
	return &Iterator[T]{
		list: list,
		curr: list.head,
	}
}

// Stream streams all elements.
//
// It is safe to modify the List while streaming.
func (list *List[T]) Stream(yield func(T) bool) {
	for curr := list.head; curr != nil; curr = curr.next {
		if !yield(curr.Value) {
			return
		}
	}
}

// Stream2 streams all elements with their indices.
//
// It is safe to modify the Array while streaming.
func (list *List[T]) Stream2(yield func(int, T) bool) {
	for curr, i := list.head, 0; curr != nil; curr, i = curr.next, i+1 {
		if !yield(i, curr.Value) {
			return
		}
	}
}

// FindIndex returns the index of the first element
// that satisfies the specified predicate.
// If no such element is found, 0 and false are returned.
func (list *List[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	for curr, i := list.head, 0; curr != nil; curr, i = curr.next, i+1 {
		if predicate(curr.Value) {
			return i, true
		}
	}

	return -1, false
}

// FindRef returns a reference to the first element
// that matches the specified predicate.
// If no element matches the predicate, nil and false are returned.
func (list *List[T]) FindRef(predicate predication.Predicate[T]) (*T, bool) {
	node, ok := list.FindNode(predicate)
	if !ok {
		return nil, false
	}

	return &node.Value, true
}

// FindNode returns the first Node
// that matches the specified predicate.
// If no node matches the predicate, nil and false are returned.
func (list *List[T]) FindNode(predicate predication.Predicate[T]) (*Node[T], bool) {
	for curr := list.head; curr != nil; curr = curr.next {
		if predicate(curr.Value) {
			return curr, true
		}
	}

	return nil, false
}

// Head returns the first Node.
func (list *List[T]) Head() *Node[T] {
	return list.head
}

// Tail returns the last Node.
func (list *List[T]) Tail() *Node[T] {
	return list.tail
}
