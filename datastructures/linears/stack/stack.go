package stack

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Stack[T any] array.Array[T]

func New[T any]() *Stack[T] { return NewWithCapacity[T](0) }

func NewWithCapacity[T any](initialCapacity int) *Stack[T] {
	arr := array.NewWithCapacity[T](initialCapacity)
	return (*Stack[T])(arr)
}

func (stack *Stack[T]) array() *array.Array[T] { return (*array.Array[T])(stack) }

func (stack *Stack[T]) Empty() bool  { return stack.array().Size() == 0 }
func (stack *Stack[T]) Push(value T) { stack.array().Append(value) }
func (stack *Stack[T]) Peek() T      { return stack.array().Get(-1) }
func (stack *Stack[T]) Pop() T {
	defer stack.array().Remove(-1)
	return stack.Peek()
}

func (stack *Stack[T]) ForEach(callback functions.ParamCallback[T]) {
	for !stack.Empty() {
		callback(stack.Pop())
	}
}
