package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

func (stream Stream[T]) ForEach(function functions.ParamCallback[T]) {
	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		function(elem.Value)
	}
}

func Reduce[T, P any](stream Stream[T], accumulator P, reducer functions.Reducer[T, P]) P {
	acc := accumulator

	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		acc = reducer(acc, elem.Value)
	}

	return acc
}

func (stream Stream[T]) Any(predictor functions.Predictor[T]) bool {
	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		if predictor(elem.Value) {
			return true
		}
	}

	return false
}

func (stream Stream[T]) All(predictor functions.Predictor[T]) bool {
	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
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

	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		count++
	}

	return count
}

func (stream Stream[T]) Max(comparator functions.Comparator[T]) optional.Optional[T] {
	var max T
	set := false

	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		if !set || comparator(elem.Value, max) == comparison.FirstBigger {
			max = elem.Value
			set = true
		}
	}

	return optional.Optional[T]{max, set}
}

func (stream Stream[T]) Min(comparator functions.Comparator[T]) optional.Optional[T] {
	var min T
	set := false

	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		if !set || comparator(elem.Value, min) == comparison.FirstSmaller {
			min = elem.Value
			set = true
		}
	}

	return optional.Optional[T]{min, set}
}

func (stream Stream[T]) First() optional.Optional[T] {
	return stream.supplier.Supply()
}

func (stream Stream[T]) Find(predictor functions.Predictor[T]) optional.Optional[T] {
	for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
		if predictor(elem.Value) {
			return optional.FromValue(elem.Value)
		}
	}

	return optional.Empty[T]()
}
