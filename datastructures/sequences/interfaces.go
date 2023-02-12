package sequences

import "github.com/djordje200179/extendedlibrary/misc/functions"

type Queue[T any] interface {
	Empty() bool

	Push(value T)
	Peek() T
	Pop() T

	ForEach(callback functions.ParamCallback[T])
}

type Stack[T any] interface {
	Empty() bool

	Push(value T)
	Peek() T
	Pop() T

	ForEach(callback functions.ParamCallback[T])
}

type Deque[T any] interface {
	Empty() bool

	PushFront(value T)
	PeekFront() T
	PopFront() T

	PushBack(value T)
	PeekBack() T
	PopBack() T

	ForEach(callback functions.ParamCallback[T])
}

type PriorityQueue[T any] interface {
	Empty() bool

	Push(value T, priority int)
	Peek() T
	Pop() T

	ForEach(callback functions.ParamCallback[T])
}