package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

func (stream Stream[T]) ForEach(function functions.ParamCallback[T]) {
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		function(elem.Value)
	}
}

type Reducer[T, P any] func(acc P, value T) P

func Reduce[T, P any](stream Stream[T], accumulator P, reducer Reducer[T, P]) P {
	acc := accumulator

	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		acc = reducer(acc, elem.Value)
	}

	return acc
}

func (stream Stream[T]) Reduce(accumulator T, reducer Reducer[T, T]) T {
	return Reduce(stream, accumulator, reducer)
}

func (stream Stream[T]) Any(predicate predication.Predicate[T]) bool {
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if predicate(elem.Value) {
			return true
		}
	}

	return false
}

func (stream Stream[T]) All(predicate predication.Predicate[T]) bool {
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if !predicate(elem.Value) {
			return false
		}
	}

	return true
}

type Collector[T, R any] interface {
	Supply(value T)
	Finish() R
}

func Collect[T, R any](stream Stream[T], collector Collector[T, R]) R {
	stream.ForEach(func(elem T) {
		collector.Supply(elem)
	})

	return collector.Finish()
}

func (stream Stream[T]) Collect(collector Collector[T, T]) T {
	return Collect(stream, collector)
}

func (stream Stream[T]) Count() int {
	count := 0
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		count++
	}

	return count
}

func (stream Stream[T]) Max(comparator comparison.Comparator[T]) optional.Optional[T] {
	var currMax T
	set := false

	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if !set || comparator(elem.Value, currMax) == comparison.FirstBigger {
			currMax = elem.Value
			set = true
		}
	}

	return optional.Optional[T]{
		Value: currMax,
		Valid: set,
	}
}

func (stream Stream[T]) Min(comparator comparison.Comparator[T]) optional.Optional[T] {
	var currMin T
	set := false

	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if !set || comparator(elem.Value, currMin) == comparison.FirstSmaller {
			currMin = elem.Value
			set = true
		}
	}

	return optional.Optional[T]{
		Value: currMin,
		Valid: set,
	}
}

func (stream Stream[T]) First() optional.Optional[T] {
	return stream.supplier()
}

func (stream Stream[T]) Find(predicate predication.Predicate[T]) optional.Optional[T] {
	return stream.Filter(predicate).First()
}

func (stream Stream[T]) Channel() <-chan T {
	channel := make(chan T)

	go func() {
		for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
			channel <- elem.Value
		}

		close(channel)
	}()

	return channel
}
