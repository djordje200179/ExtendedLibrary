package streams

type iterator[T any] struct {
	stream  *Stream[T]
	current T
	
	started, ended bool
}

func (it *iterator[T]) Valid() bool { return !it.ended }

func (it *iterator[T]) Move() {
	if elem := it.stream.getNext(); elem.HasValue() {
		it.current = elem.Get()
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
