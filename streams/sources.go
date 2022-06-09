package streams

import (
	"github.com/djordje200179/GoExtendedLibrary/misc/functions"
)

func Supply[T any](supplier func() T) Stream[T] {
	stream := create[T]()

	go func() {
		for stream.waitRequest() {
			stream.data <- supplier()
		}

		stream.close()
	}()

	return stream
}

func Generate[T any](seed T, generator func(curr T) T) Stream[T] {
	stream := create[T]()

	go func() {
		for curr := seed; stream.waitRequest(); curr = generator(curr) {
			stream.data <- curr
		}

		stream.close()
	}()

	return stream
}

func FromValues[T any](values ...T) Stream[T] {
	stream := create[T]()

	go func() {
		for i := 0; i < len(values) && stream.waitRequest(); i++ {
			stream.data <- values[i]
		}

		stream.close()
	}()

	return stream
}

func FromChannel[T any](ch <-chan T) Stream[T] {
	stream := create[T]()

	go func() {
		for stream.waitRequest() {
			data, ok := <-ch

			if !ok {
				break
			}

			stream.data <- data
		}

		stream.close()
	}()

	return stream
}

func Range(lower, upper int) Stream[int] {
	return Generate(lower, functions.Increment).Limit(upper - lower)
}