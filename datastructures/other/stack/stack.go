package stack

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Stack[T any] *linkedlist.LinkedList[T]

func New[T any]() Stack[T] {
	return linkedlist.New[T]()
}

func (stack Stack[T]) getList() *linkedlist.LinkedList[T] {
	return stack
}

func (stack Stack[T]) Push(value T) {
	stack.getList().Append(value)
}

func (stack Stack[T]) Pop() T {
	return stack.getList().Remove(-1)
}

func (stack Stack[T]) Peek() T {
	return stack.getList().Get(-1)
}

func (stack Stack[T]) IsEmpty() bool {
	return stack.getList().Size() == 0
}
