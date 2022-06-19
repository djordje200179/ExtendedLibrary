package streams

type iterator[T any] struct {
	stream  *Stream[T]
	current T
	started bool
	ended   bool
}

func (it *iterator[T]) Valid() bool {
	return !it.ended
}

func (it *iterator[T]) Move() {
	data, ok := it.stream.getNext().GetPair()
	if ok {
		it.current = data
	} else {
		it.ended = true
	}
}

func (it *iterator[T]) Get() T {
	if !it.started {
		it.Move()
		it.started = true
	}

	return it.current
}
