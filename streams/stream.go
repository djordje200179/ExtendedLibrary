package streams

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
)

type Stream[T any] struct {
	supplier suppliers.Supplier[T]
}

type Streamer[T any] interface {
	Stream() Stream[T]
}

func FromChannel[T any](channel <-chan T) Stream[T] { return Stream[T]{suppliers.Channel[T]{channel}} }

func FromValues[T any](values ...T) Stream[T] { return FromSlice(values) }
func FromRange(lower, upper int) Stream[int]  { return Stream[int]{suppliers.Range(lower, upper)} }

func FromSlice[T any](slice []T) Stream[T]      { return Stream[T]{suppliers.Slice(slice)} }
func FromSliceRefs[T any](slice []T) Stream[*T] { return Stream[*T]{suppliers.SliceRefs(slice)} }

func FromMap[K comparable, V any](m map[K]V) Stream[misc.Pair[K, V]] {
	return Stream[misc.Pair[K, V]]{suppliers.Map(m)}
}

func FromFiniteGenerator[T any](generator functions.EmptyGenerator[optional.Optional[T]]) Stream[T] {
	return Stream[T]{suppliers.Function[T](generator)}
}

func FromInfiniteGenerator[T any](generator functions.EmptyGenerator[T]) Stream[T] {
	return Stream[T]{suppliers.Infinite(generator)}
}
