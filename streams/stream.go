package streams

import "github.com/djordje200179/GoExtendedLibrary/misc/optional"

type signal bool

const (
	end  signal = false
	next signal = true
)

type Stream[T any] struct {
	data   chan T
	signal chan signal
}

func create[T any]() Stream[T] {
	return Stream[T]{
		data:   make(chan T),
		signal: make(chan signal),
	}
}

func (stream Stream[T]) close() {
	close(stream.data)
	close(stream.signal)
}

func (stream Stream[T]) getNext() optional.Optional[T] {
	stream.signal <- next

	data, ok := <-stream.data
	return optional.New(data, ok)
}

func (stream Stream[T]) stop() {
	stream.signal <- end
}

func (stream Stream[T]) waitRequest() bool {
	return <-stream.signal == next
}

type Streamer[T any] interface {
	Stream() Stream[T]
}
