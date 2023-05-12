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

func (stream Stream[T]) Any(predictor predication.Predictor[T]) bool {
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if predictor(elem.Value) {
			return true
		}
	}

	return false
}

func (stream Stream[T]) All(predictor predication.Predictor[T]) bool {
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if !predictor(elem.Value) {
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

func (stream Stream[T]) Count() int {
	count := 0
	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		count++
	}

	return count
}

func (stream Stream[T]) Max(comparator comparison.Comparator[T]) optional.Optional[T] {
	var max T
	set := false

	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if !set || comparator(elem.Value, max) == comparison.FirstBigger {
			max = elem.Value
			set = true
		}
	}

	return optional.Optional[T]{
		Value: max,
		Valid: set,
	}
}

func (stream Stream[T]) Min(comparator comparison.Comparator[T]) optional.Optional[T] {
	var min T
	set := false

	for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
		if !set || comparator(elem.Value, min) == comparison.FirstSmaller {
			min = elem.Value
			set = true
		}
	}

	return optional.Optional[T]{
		Value: min,
		Valid: set,
	}
}

func (stream Stream[T]) First() optional.Optional[T] {
	return stream.supplier()
}

func (stream Stream[T]) Find(predictor predication.Predictor[T]) optional.Optional[T] {
	return stream.Filter(predictor).First()
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
