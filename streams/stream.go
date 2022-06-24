package streams

import (
	"github.com/djordje200179/extendedlibrary/concurrency/messenger"
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type signal bool

const end, next signal = false, true

type Stream[T any] struct {
	dataChannel chan T
	signaler    messenger.Messenger[signal]

	closed bool
}

func create[T any]() *Stream[T] {
	stream := new(Stream[T])

	stream.dataChannel = make(chan T)
	stream.signaler = messenger.New[signal](1)

	return stream
}

func (stream *Stream[T]) close() {
	stream.closed = true
	close(stream.dataChannel)
	stream.signaler.Close()
}

func (stream *Stream[T]) getNext() optional.Optional[T] {
	if stream.closed {
		return optional.Empty[T]()
	}

	stream.signaler.Send(next)

	data, ok := <-stream.dataChannel
	return optional.New(data, ok)
}

func (stream *Stream[T]) stop() {
	if !stream.closed {
		stream.signaler.Send(end)
	}
}

func (stream *Stream[T]) waitRequest() bool { return stream.signaler.ReadSync().Get() == next }

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
