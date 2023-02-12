package sequences

type Queue[T any] interface {
	Empty() bool

	PushBack(value T)
	PopFront() T
	PeekFront() T
}

type Stack[T any] interface {
	Empty() bool

	PushBack(value T)
	PopBack() T
	PeekBack() T
}

type Deque[T any] interface {
	Queue[T]
	Stack[T]
}
