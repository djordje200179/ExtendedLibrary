package seqs

type Queue[T any] interface {
	Empty() bool

	PushBack(value T)
	PopFront() T
}

type PeekableQueue[T any] interface {
	Queue[T]

	PeekFront() T
}

type Stack[T any] interface {
	Empty() bool

	PushBack(value T)
	PopBack()
}

type PeekableStack[T any] interface {
	Stack[T]

	PeekBack() T
}

type Deque[T any] interface {
	Queue[T]
	Stack[T]
}

type PeekableDeque[T any] interface {
	Deque[T]

	PeekFront() T
	PeekBack() T
}
