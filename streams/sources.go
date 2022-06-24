package streams

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

func Supply[T any](supplier functions.EmptyGenerator[T]) *Stream[T] {
	stream := create[T]()

	go func() {
		for stream.waitRequest() {
			stream.dataChannel <- supplier()
		}

		stream.close()
	}()

	return stream
}

func Generate[T any](seed T, generator functions.ParamGenerator[T, T]) *Stream[T] {
	stream := create[T]()

	go func() {
		for curr := seed; stream.waitRequest(); curr = generator(curr) {
			stream.dataChannel <- curr
		}

		stream.close()
	}()

	return stream
}

func FromValues[T any](values ...T) *Stream[T] {
	return FromSlice[T](values)
}

func FromSlice[T any](values []T) *Stream[T] {
	stream := create[T]()

	go func() {
		for i := 0; i < len(values) && stream.waitRequest(); i++ {
			stream.dataChannel <- values[i]
		}

		stream.close()
	}()

	return stream
}

func FromChannel[T any](ch <-chan T) *Stream[T] {
	stream := create[T]()

	go func() {
		for stream.waitRequest() {
			data, ok := <-ch

			if !ok {
				break
			}

			stream.dataChannel <- data
		}

		stream.close()
	}()

	return stream
}

func FromIterable[T any](iterable datastructures.Iterable[T]) *Stream[T] {
	stream := create[T]()

	go func() {
		for it := iterable.Iterator(); it.Valid() && stream.waitRequest(); it.Move() {
			stream.dataChannel <- it.Get()
		}

		stream.close()
	}()

	return stream
}

func Range(lower, upper int) *Stream[int] {
	increment := func(curr int) int { return curr + 1 }
	return Generate(lower, increment).Limit(upper - lower)
}
