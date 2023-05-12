package collectors

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

type partitionCollector[T any] struct {
	falseElements, trueElements []T

	predictor predication.Predictor[T]
}

func Partition[T any](predictor predication.Predictor[T]) streams.Collector[T, misc.Pair[[]T, []T]] {
	return partitionCollector[T]{predictor: predictor}
}

func (collector partitionCollector[T]) Supply(value T) {
	if collector.predictor(value) {
		collector.trueElements = append(collector.trueElements, value)
	} else {
		collector.falseElements = append(collector.falseElements, value)
	}
}

func (collector partitionCollector[T]) Finish() misc.Pair[[]T, []T] {
	return misc.Pair[[]T, []T]{collector.falseElements, collector.trueElements}
}
