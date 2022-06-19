package streams

import (
	"github.com/djordje200179/extendedlibrary/concurrency/messenger"
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type signal bool

const (
	end  signal = false
	next signal = true
)

type Stream[T any] struct {
	data     chan T
	signaler messenger.Messenger[signal]
}

func create[T any]() *Stream[T] {
	return &Stream[T]{
		data:     make(chan T),
		signaler: *messenger.New[signal](),
	}
}

func (stream *Stream[T]) close() {
	close(stream.data)
	stream.signaler.Close()
}

func (stream *Stream[T]) getNext() optional.Optional[T] {
	if stream.signaler.Closed() {
		return optional.Empty[T]()
	}

	stream.signaler.Send(next)

	data, ok := <-stream.data
	return optional.New(data, ok)
}

func (stream *Stream[T]) stop() {
	stream.signaler.Send(end)
}

func (stream *Stream[T]) waitRequest() bool {
	return stream.signaler.ReadSync().Get() == next
}

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
