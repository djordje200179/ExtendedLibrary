package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

func (stream Stream[T]) ForEach(function functions.ParamCallback[T]) {
	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		function(elem.Get())
	}
}

func Reduce[T, P any](stream Stream[T], accumulator P, reducer functions.Reducer[T, P]) P {
	acc := accumulator

	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		acc = reducer(acc, elem.Get())
	}

	return acc
}

func (stream Stream[T]) ReduceWithAny(accumulator any, reducer functions.Reducer[T, any]) any {
	return Reduce[T, any](stream, accumulator, reducer)
}

func (stream Stream[T]) Any(predictor functions.Predictor[T]) bool {
	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		if predictor(elem.Get()) {
			stream.stop()
			return true
		}
	}

	return false
}

func Collect[T, R any](stream Stream[T], collector Collector[T, R]) R {
	stream.ForEach(func(elem T) {
		collector.Supply(elem)
	})

	return collector.Finish()
}

func (stream Stream[T]) CollectWithAny(collector Collector[T, any]) any {
	return Collect[T, any](stream, collector)
}

func (stream Stream[T]) All(predictor functions.Predictor[T]) bool {
	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		if !predictor(elem.Get()) {
			stream.stop()
			return false
		}
	}

	return true
}

func (stream Stream[T]) Count() int {
	count := 0

	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		count++
	}

	return count
}

func (stream Stream[T]) Max(comparator functions.Comparator[T]) optional.Optional[T] {
	var max T
	set := false

	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		data := elem.Get()
		if !set || comparator(data, max) == comparison.FirstBigger {
			max = data
			set = true
		}
	}

	return optional.New(max, set)
}

func (stream Stream[T]) Min(comparator functions.Comparator[T]) optional.Optional[T] {
	var min T
	set := false

	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		data := elem.Get()
		if !set || comparator(data, min) == comparison.FirstSmaller {
			min = data
			set = true
		}
	}

	return optional.New(min, set)
}

func (stream Stream[T]) First(predictor functions.Predictor[T]) optional.Optional[T] {
	for {
		elem := stream.getNext()
		if !elem.HasValue() {
			break
		}

		value := elem.Get()
		if predictor(value) {
			stream.stop()
			return optional.FromValue(value)
		}
	}

	return optional.Empty[T]()
}
