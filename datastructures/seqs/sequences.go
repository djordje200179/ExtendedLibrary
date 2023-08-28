package seqs

// Queue is a first-in-first-out data structure.
type Queue[T any] interface {
	Empty() bool // Empty returns true if the queue is empty.

	PushBack(value T) // PushBack adds a value to the back of the queue.
	PopFront() T      // PopFront removes and returns the value at the front of the queue.
}

// PeekableQueue is a queue that allows peeking at the front value.
type PeekableQueue[T any] interface {
	Queue[T]

	PeekFront() T // PeekFront returns the value at the front of the queue without removing it.
}

// Stack is a last-in-first-out data structure.
type Stack[T any] interface {
	Empty() bool // Empty returns true if the stack is empty.

	PushBack(value T) // PushBack adds a value to the back of the stack.
	PopBack()         // PopBack removes the value at the back of the stack.
}

// PeekableStack is a stack that allows peeking at the back value.
type PeekableStack[T any] interface {
	Stack[T]

	PeekBack() T // PeekBack returns the value at the back of the stack without removing it.
}

// Deque is a double-ended queue.
// Values can be added and removed from both the front and back.
type Deque[T any] interface {
	Queue[T]
	Stack[T]
}

// PeekableDeque is a deque that allows peeking at both the front and back values.
type PeekableDeque[T any] interface {
	PeekableQueue[T]
	PeekableStack[T]
}
