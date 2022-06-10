package deque

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Deque[T any] *linkedlist.LinkedList[T]

func New[T any]() Deque[T] {
	return linkedlist.New[T]()
}

func (deque Deque[T]) getList() *linkedlist.LinkedList[T] {
	return deque
}

func (deque Deque[T]) PushFront(value T) {
	deque.getList().Insert(0, value)
}

func (deque Deque[T]) PushBack(value T) {
	deque.getList().Append(value)
}

func (deque Deque[T]) PopFront() T {
	return deque.getList().Remove(0)
}

func (deque Deque[T]) PopBack() T {
	return deque.getList().Remove(-1)
}

func (deque Deque[T]) PeekFront() T {
	return deque.getList().Get(0)
}

func (deque Deque[T]) PeekBack() T {
	return deque.getList().Get(-1)
}

func (deque Deque[T]) IsEmpty() bool {
	return deque.getList().Size() == 0
}
