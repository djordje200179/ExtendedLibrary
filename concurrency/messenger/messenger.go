package messenger

import "github.com/djordje200179/extendedlibrary/misc/optional"

type Messenger[T any] struct {
	ch     chan T
	closed bool
}

func New[T any]() *Messenger[T] {
	return &Messenger[T]{
		ch:     make(chan T),
		closed: false,
	}
}

func (messenger *Messenger[T]) Send(signal T) {
	messenger.ch <- signal
}

func (messenger *Messenger[T]) Read() optional.Optional[T] {
	if messenger.closed {
		return optional.Empty[T]()
	} else {
		return optional.FromValue(<-messenger.ch)
	}
}

func (messenger *Messenger[T]) IsClosed() bool {
	return messenger.closed
}

func (messenger *Messenger[T]) Close() {
	messenger.closed = true
	close(messenger.ch)
}
