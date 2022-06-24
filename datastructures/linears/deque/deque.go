package deque

import "github.com/djordje200179/extendedlibrary/datastructures/sequences/array"

type Deque[T any] array.Array[T]

func New[T any]() *Deque[T] { return NewWithCapacity[T](0) }

func NewWithCapacity[T any](initialCapacity int) *Deque[T] {
	arr := array.NewWithCapacity[T](initialCapacity)
	return (*Deque[T])(arr)
}

func (deque *Deque[T]) array() *array.Array[T] { return (*array.Array[T])(deque) }

func (deque *Deque[T]) Empty() bool { return deque.array().Size() == 0 }

func (deque *Deque[T]) PushFront(value T) { deque.array().Insert(0, value) }
func (deque *Deque[T]) PushBack(value T)  { deque.array().Append(value) }

func (deque *Deque[T]) PeekFront() T { return deque.array().Get(0) }
func (deque *Deque[T]) PeekBack() T  { return deque.array().Get(-1) }

func (deque *Deque[T]) PopFront() T {
	defer deque.array().Remove(0)
	return deque.PeekFront()
}

func (deque *Deque[T]) PopBack() T {
	defer deque.array().Remove(-1)
	return deque.PeekBack()
}
