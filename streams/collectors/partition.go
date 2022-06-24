package collectors

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type partitionCollector[T any] struct {
	falsy, truly sequences.Sequence[T]
	predictor    functions.Predictor[T]
}

func Partition[T any](predictor functions.Predictor[T]) streams.Collector[T, misc.Pair[sequences.Sequence[T], sequences.Sequence[T]]] {
	return partitionCollector[T]{
		falsy:     array.New[T](),
		truly:     array.New[T](),
		predictor: predictor,
	}
}

func (collector partitionCollector[T]) Supply(value T) {
	var seq sequences.Sequence[T]
	if collector.predictor(value) {
		seq = collector.truly
	} else {
		seq = collector.falsy
	}
	seq.Append(value)
}

func (collector partitionCollector[T]) Finish() misc.Pair[sequences.Sequence[T], sequences.Sequence[T]] {
	return misc.Pair[sequences.Sequence[T], sequences.Sequence[T]]{collector.falsy, collector.truly}
}
