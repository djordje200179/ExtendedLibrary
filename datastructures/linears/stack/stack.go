package stack

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Stack[T any] struct {
	sequence sequences.Sequence[T]
}

func NewFrom[T any](sequence sequences.Sequence[T]) Stack[T] { return Stack[T]{sequence} }
func New[T any]() Stack[T]                                   { return Stack[T]{array.New[T]()} }

func (stack Stack[T]) Empty() bool { return stack.sequence.Size() == 0 }

func (stack Stack[T]) Push(value T) { stack.sequence.Append(value) }
func (stack Stack[T]) Peek() T      { return stack.sequence.Get(-1) }
func (stack Stack[T]) Pop() T {
	defer stack.sequence.Remove(-1)
	return stack.Peek()
}

func (stack Stack[T]) ForEach(callback functions.ParamCallback[T]) {
	for !stack.Empty() {
		callback(stack.Pop())
	}
}
