package stack

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

type Stack[T any] struct {
	seq sequences.Sequence[T]
}

func New[T any]() Stack[T] {
	return Stack[T]{linkedlist.New[T]()}
}

func (stack Stack[T]) Push(value T) {
	stack.seq.Append(value)
}

func (stack Stack[T]) Pop() T {
	defer stack.seq.Remove(-1)
	return stack.Peek()
}

func (stack Stack[T]) Peek() T {
	return stack.seq.Get(-1)
}

func (stack Stack[T]) IsEmpty() bool {
	return stack.seq.Size() == 0
}
