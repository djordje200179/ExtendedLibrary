package collectors

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type partitionCollector[T any] struct {
	falsy, truly []T
	predictor    functions.Predictor[T]
}

func Partition[T any](predictor functions.Predictor[T]) streams.Collector[T, misc.Pair[[]T, []T]] {
	return partitionCollector[T]{predictor: predictor}
}

func (collector partitionCollector[T]) Supply(value T) {
	if collector.predictor(value) {
		collector.truly = append(collector.truly, value)
	} else {
		collector.falsy = append(collector.falsy, value)
	}
}

func (collector partitionCollector[T]) Finish() misc.Pair[[]T, []T] {
	return misc.Pair[[]T, []T]{collector.falsy, collector.truly}
}
