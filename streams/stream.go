package streams

import (
	"bufio"
	"github.com/djordje200179/extendedlibrary/misc/math"
	"io"
	"iter"
)

type Stream[T any] iter.Seq[T]

type Streamer[T any] interface {
	Stream(yield func(T) bool)
}

func From[T any](streamer Streamer[T]) Stream[T] {
	return streamer.Stream
}

func FromSlice[T any](slice []T) Stream[T] {
	return func(yield func(T) bool) {
		for _, elem := range slice {
			if !yield(elem) {
				break
			}
		}
	}
}

func FromChannel[T any](ch <-chan T) Stream[T] {
	return func(yield func(T) bool) {
		for elem := range ch {
			if !yield(elem) {
				break
			}
		}
	}
}

func FromReader(reader io.Reader, splitFunc bufio.SplitFunc) Stream[string] {
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitFunc)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
	}
}

func FromGenerator[T any](generator func() T) Stream[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(generator()) {
				break
			}
		}
	}
}

func RangeIncrement[T math.Real](lower, upper, increment T) Stream[T] {
	return func(yield func(T) bool) {
		for i := lower; i < upper; i += increment {
			if !yield(i) {
				break
			}
		}
	}
}

func Range[T math.Real](lower, upper T) Stream[T] {
	return RangeIncrement(lower, upper, 1)
}

func Enumerate[T any](s Stream[T]) Stream2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for elem := range s {
			if !yield(i, elem) {
				break
			}

			i++
		}
	}
}

func Chunk[T any](s Stream[T], size int) Stream[[]T] {
	return func(yield func([]T) bool) {
		chunk := make([]T, 0, size)

		for elem := range s {
			chunk = append(chunk, elem)
			if len(chunk) == size {
				if !yield(chunk) {
					break
				}

				chunk = make([]T, 0, size)
			}
		}
	}
}

func Window[T any](s Stream[T], width int) Stream[[]T] {
	return func(yield func([]T) bool) {
		window := make([]T, width)

		i := 0
		for elem := range s {
			if i < width {
				window[i] = elem
				i++
				continue
			}

			if !yield(window) {
				break
			}

			for i := range width - 1 {
				window[i] = window[i+1]
			}

			window[width-1] = elem
		}
	}
}
