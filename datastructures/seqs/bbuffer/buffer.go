package bbuffer

import (
	"github.com/djordje200179/extendedlibrary/datastructures/seqs"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Buffer[T any] chan T

func New[T any](size int) Buffer[T] {
	return make(chan T, size)
}

func FromChannel[T any](channel chan T) Buffer[T] {
	return channel
}

func Collector[T any](size int) streams.Collector[T, seqs.Queue[T]] {
	return seqs.Collector[T, seqs.Queue[T]]{New[T](size)}
}

func (buffer Buffer[T]) Empty() bool {
	return len(buffer) == 0
}

func (buffer Buffer[T]) PushBack(value T) {
	buffer <- value
}

func (buffer Buffer[T]) PopFront() T {
	return <-buffer
}
