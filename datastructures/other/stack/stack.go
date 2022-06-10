package stack

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Stack[T any] struct {
	list *linkedlist.LinkedList[T]
}

func New[T any]() Stack[T] {
	return Stack[T]{linkedlist.New[T]()}
}

func (stack Stack[T]) Push(value T) {
	stack.list.Append(value)
}

func (stack Stack[T]) Pop() T {
	return stack.list.Remove(-1)
}

func (stack Stack[T]) Peek() T {
	return stack.list.Get(-1)
}

func (stack Stack[T]) IsEmpty() bool {
	return stack.list.Size() == 0
}
