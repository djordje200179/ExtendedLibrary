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

func New[T any](supplier suppliers.Supplier[T]) Stream[T] { return Stream[T]{supplier} }
func FromChannel[T any](channel <-chan T) Stream[T]       { return New(suppliers.FromChannel(channel)) }
func FromRange(lower, upper int) Stream[int]              { return New(suppliers.FromRange(lower, upper)) }

func FromSlice[T any](slice []T) Stream[T]      { return New(suppliers.FromSlice(slice)) }
func FromSliceRefs[T any](slice []T) Stream[*T] { return New(suppliers.FromSliceRefs(slice)) }
func FromValues[T any](values ...T) Stream[T]   { return FromSlice(values) }

func FromMap[K comparable, V any](m map[K]V) Stream[misc.Pair[K, V]] {
	return New(suppliers.FromMap(m))
}

func FromFiniteGenerator[T any](generator functions.EmptyGenerator[optional.Optional[T]]) Stream[T] {
	return New[T](suppliers.FunctionSupplier[T](generator))
}

func FromInfiniteGenerator[T any](generator functions.EmptyGenerator[T]) Stream[T] {
	return New(suppliers.InfiniteGenerator(generator))
}
