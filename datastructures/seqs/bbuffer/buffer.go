package bbuffer

// Buffer is a seqs.Queue with a fixed size.
// Operations are blocking if the buffer is full or empty.
//
// The buffer is implemented using a channel,
// so it is thread-safe and can be used in concurrent programs.
type Buffer[T any] chan T

// New creates a new Buffer with the given size.
func New[T any](size int) Buffer[T] { return make(chan T, size) }

// NewUnbuffered creates a new Buffer with size 0.
func NewUnbuffered[T any]() Buffer[T] { return make(chan T) }

// FromChannel creates a Buffer from the given channel.
func FromChannel[T any](channel chan T) Buffer[T] { return channel }

// Empty returns true if the Buffer is empty.
func (b Buffer[T]) Empty() bool { return len(b) == 0 }

// PushBack adds the given value to the back.
// If the Buffer is full, the operation blocks
// until there is space available.
func (b Buffer[T]) PushBack(value T) { b <- value }

// TryPushBack tries to add the given value to the back
// and returns true if successful.
func (b Buffer[T]) TryPushBack(value T) bool {
	select {
	case b <- value:
		return true
	default:
		return false
	}
}

// PopFront removes and returns the value at the front.
// If the Buffer is empty, the operation blocks
// until there is a value available.
func (b Buffer[T]) PopFront() T { return <-b }

// TryPopFront tries to remove and return
// the value at the front and true if successful.
func (b Buffer[T]) TryPopFront() (T, bool) {
	select {
	case value := <-b:
		return value, true
	default:
		var zero T
		return zero, false
	}
}
