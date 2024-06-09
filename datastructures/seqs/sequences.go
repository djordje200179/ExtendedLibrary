package seqs

// BackPusher allows adding values to the back.
type BackPusher[T any] interface {
	// PushBack adds the given value to the back.
	PushBack(value T)

	// TryPushBack tries to add the given value to the back
	// and returns true if successful.
	TryPushBack(value T) bool
}

// FrontPusher allows adding values to the front.
type FrontPusher[T any] interface {
	// PushFront adds the given value to the front.
	PushFront(value T)

	// TryPushFront tries to add the given value to the front
	// and returns true if successful.
	TryPushFront(value T) bool
}

// BackPopper allows removing values from the back.
type BackPopper[T any] interface {
	// PopBack removes and returns the value at the back.
	PopBack()

	// TryPopBack tries to remove and return
	// the value at the back and true if successful.
	TryPopBack() (T, bool)
}

// FrontPopper allows removing values from the front.
type FrontPopper[T any] interface {
	// PopFront removes and returns the value at the front.
	PopFront()

	// TryPopFront tries to remove and return
	// the value at the front and true if successful.
	TryPopFront() (T, bool)
}

// BackPeeker allows peeking at the back value.
type BackPeeker[T any] interface {
	// PeekBack returns the value at the back
	// without removing it.
	PeekBack() T

	// TryPeekBack returns the value at the back
	// without removing it and true if successful.
	TryPeekBack() (T, bool)
}

// FrontPeeker allows peeking at the front value.
type FrontPeeker[T any] interface {
	// PeekFront returns the value at the front
	// without removing it.
	PeekFront() T

	// TryPeekFront returns the value at the front
	// without removing it and true if successful.
	TryPeekFront() (T, bool)
}

// Queue is a first-in-first-out data structure.
type Queue[T any] interface {
	// Empty returns true if there are no elements.
	Empty() bool

	BackPusher[T]
	FrontPopper[T]
}

// PeekableQueue is a Queue that allows peeking at the front value.
type PeekableQueue[T any] interface {
	Queue[T]
	FrontPeeker[T]
}

// Stack is a last-in-first-out data structure.
type Stack[T any] interface {
	// Empty returns true if there are no elements.
	Empty() bool

	BackPusher[T]
	BackPopper[T]
}

// PeekableStack is a Stack that allows peeking at the back value.
type PeekableStack[T any] interface {
	Stack[T]
	BackPeeker[T]
}

// Deque is a double-ended Queue.
// Therefore, it is both a Queue and a Stack.
// Values can be added and removed from both the front and back.
type Deque[T any] interface {
	Queue[T]
	Stack[T]
}

// PeekableDeque is a Deque that allows peeking
// at both the front and back values.
type PeekableDeque[T any] interface {
	PeekableQueue[T]
	PeekableStack[T]
}
