package signaler

import "github.com/djordje200179/extendedlibrary/misc/optional"

type Signaler[T any] struct {
	ch     chan T
	closed bool
}

func New[T any]() *Signaler[T] {
	return &Signaler[T]{
		ch:     make(chan T),
		closed: false,
	}
}

func (signaler *Signaler[T]) Send(signal T) {
	signaler.ch <- signal
}

func (signaler *Signaler[T]) Read() optional.Optional[T] {
	if signaler.closed {
		return optional.Empty[T]()
	} else {
		return optional.FromValue(<-signaler.ch)
	}
}

func (signaler *Signaler[T]) IsClosed() bool {
	return signaler.closed
}

func (signaler *Signaler[T]) Close() {
	signaler.closed = true
	close(signaler.ch)
}
