package deque

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Deque[T any] struct {
	list *linkedlist.LinkedList[T]
}

func New[T any]() Deque[T] {
	return Deque[T]{linkedlist.New[T]()}
}

func (deque Deque[T]) PushFront(value T) {
	deque.list.Insert(0, value)
}

func (deque Deque[T]) PushBack(value T) {
	deque.list.Append(value)
}

func (deque Deque[T]) PopFront() T {
	defer deque.list.Remove(0)
	return deque.PeekFront()
}

func (deque Deque[T]) PopBack() T {
	defer deque.list.Remove(-1)
	return deque.PeekBack()
}

func (deque Deque[T]) PeekFront() T {
	return deque.list.Get(0)
}

func (deque Deque[T]) PeekBack() T {
	return deque.list.Get(-1)
}

func (deque Deque[T]) IsEmpty() bool {
	return deque.list.Size() == 0
}
