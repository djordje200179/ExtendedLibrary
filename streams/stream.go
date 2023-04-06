package streams

import (
	"bufio"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"io"
)

type Stream[T any] struct {
	suppliers.Supplier[T]
}

type Streamer[T any] interface {
	Stream() Stream[T]
}

func FromChannel[T any](channel <-chan T) Stream[T] {
	return Stream[T]{
		Supplier: suppliers.Channel[T](channel),
	}
}

func FromValues[T any](values ...T) Stream[T] {
	return FromSlice(values)
}

func FromRange(lower, upper int) Stream[int] {
	return Stream[int]{
		Supplier: suppliers.Range(lower, upper),
	}
}

func FromSlice[T any](slice []T) Stream[T] {
	return Stream[T]{
		Supplier: suppliers.Slice(slice),
	}
}

func FromSliceRefs[T any](slice []T) Stream[*T] {
	return Stream[*T]{
		Supplier: suppliers.SliceRefs(slice),
	}
}

func FromMap[K comparable, V any](m map[K]V) Stream[misc.Pair[K, V]] {
	return Stream[misc.Pair[K, V]]{
		Supplier: suppliers.Map(m),
	}
}

func FromFiniteGenerator[T any](generator functions.EmptyGenerator[optional.Optional[T]]) Stream[T] {
	return Stream[T]{
		Supplier: suppliers.Function[T](generator),
	}
}

func FromInfiniteGenerator[T any](generator functions.EmptyGenerator[T]) Stream[T] {
	return Stream[T]{
		Supplier: suppliers.Infinite(generator),
	}
}

func FromReader(reader io.Reader, splitFunction bufio.SplitFunc) Stream[string] {
	return Stream[string]{
		Supplier: suppliers.Reader(reader, splitFunction),
	}
}
