package collectors

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

type partitionCollector[T any] struct {
	falseElements, trueElements []T

	predicate predication.Predicate[T]
}

func Partition[T any](predicate predication.Predicate[T]) streams.Collector[T, misc.Pair[[]T, []T]] {
	return partitionCollector[T]{predicate: predicate}
}

func (collector partitionCollector[T]) Supply(value T) {
	if collector.predicate(value) {
		collector.trueElements = append(collector.trueElements, value)
	} else {
		collector.falseElements = append(collector.falseElements, value)
	}
}

func (collector partitionCollector[T]) Finish() misc.Pair[[]T, []T] {
	return misc.MakePair(collector.falseElements, collector.trueElements)
}
