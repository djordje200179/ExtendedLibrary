package deque

type Deque[T any] struct {
	slice []T
}

func New[T any]() *Deque[T] {
	return NewWithCapacity[T](0)
}

func NewWithCapacity[T any](initialCapacity int) *Deque[T] {
	deque := new(Deque[T])
	deque.slice = make([]T, 0, initialCapacity)

	return deque
}

func (deque *Deque[T]) PushFront(value T) {
	newSlice := make([]T, len(deque.slice)+1)
	newSlice[0] = value
	copy(newSlice[1:], deque.slice)

	deque.slice = newSlice
}

func (deque *Deque[T]) PushBack(value T) {
	deque.slice = append(deque.slice, value)
}

func (deque *Deque[T]) PopFront() T {
	value := deque.PeekFront()
	deque.slice = deque.slice[1:]
	return value
}

func (deque *Deque[T]) PopBack() T {
	value := deque.PeekBack()
	deque.slice = deque.slice[:len(deque.slice)-1]
	return value
}

func (deque *Deque[T]) PeekFront() T {
	return deque.slice[0]
}

func (deque *Deque[T]) PeekBack() T {
	return deque.slice[len(deque.slice)-1]
}

func (deque *Deque[T]) Empty() bool {
	return len(deque.slice) == 0
}
