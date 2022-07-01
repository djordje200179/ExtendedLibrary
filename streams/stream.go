package streams

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type signal bool

const end, next signal = false, true

type Stream[T any] struct {
	data    chan T
	signals chan signal

	closed bool
}

func create[T any]() *Stream[T] {
	stream := new(Stream[T])

	stream.data = make(chan T)
	stream.signals = make(chan signal, 1)

	return stream
}

func (stream *Stream[T]) close() {
	stream.closed = true
	close(stream.data)
	close(stream.signals)
}

func (stream *Stream[T]) getNext() optional.Optional[T] {
	if stream.closed {
		return optional.Empty[T]()
	}

	stream.signals <- next

	data, ok := <-stream.data
	return optional.New[T](data, ok)
}

func (stream *Stream[T]) stop() {
	if !stream.closed {
		stream.signals <- end
	}
}

func (stream *Stream[T]) waitRequest() bool { return <-stream.signals == next }

func (stream *Stream[T]) Iterator() datastructures.Iterator[T] {
	return &iterator[T]{
		stream:  stream,
		started: false,
		ended:   false,
	}
}

type Streamer[T any] interface {
	Stream() *Stream[T]
}
