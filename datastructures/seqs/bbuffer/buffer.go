package bbuffer

// Buffer is a bounded buffer, which is a queue with a fixed size.
// Operations are blocking if the buffer is full or empty.
type Buffer[T any] chan T

// New creates a new bounded buffer with the given size.
func New[T any](size int) Buffer[T] {
	return make(chan T, size)
}

// FromChannel creates a bounded buffer from the given channel.
func FromChannel[T any](channel chan T) Buffer[T] {
	return channel
}

// Empty returns true if the buffer is empty.
func (buffer Buffer[T]) Empty() bool {
	return len(buffer) == 0
}

// PushBack adds the given value to the back of the buffer.
// If the buffer is full, the operation blocks until there is space available.
func (buffer Buffer[T]) PushBack(value T) {
	buffer <- value
}

// PopFront removes and returns the value at the front of the buffer.
// If the buffer is empty, the operation blocks until there is a value available.
func (buffer Buffer[T]) PopFront() T {
	return <-buffer
}
