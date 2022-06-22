package stack

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type Stack[T any] struct {
	slice []T
}

func New[T any]() *Stack[T] {
	stack := new(Stack[T])
	stack.slice = make([]T, 0)

	return stack
}

func (stack *Stack[T]) Push(value T) {
	stack.slice = append(stack.slice, value)
}

func (stack *Stack[T]) Pop() T {
	value := stack.Peek()
	stack.slice = stack.slice[:len(stack.slice)-1]
	return value
}

func (stack *Stack[T]) Peek() T {
	return stack.slice[len(stack.slice)-1]
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.slice) == 0
}

func (stack *Stack[T]) ForEach(callback functions.ParamCallback[T]) {
	for !stack.Empty() {
		callback(stack.Pop())
	}
}
