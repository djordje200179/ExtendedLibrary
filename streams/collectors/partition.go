package collectors

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

type partitionCollector[T any] struct {
	falsy, truly []T

	predicate predication.Predicate[T]
}

func Partition[T any](predicate predication.Predicate[T]) streams.Collector[T, misc.Pair[[]T, []T]] {
	return &partitionCollector[T]{predicate: predicate}
}

func (c *partitionCollector[T]) Supply(value T) {
	if c.predicate(value) {
		c.truly = append(c.truly, value)
	} else {
		c.falsy = append(c.falsy, value)
	}
}

func (c *partitionCollector[T]) Finish() misc.Pair[[]T, []T] {
	return misc.MakePair(c.falsy, c.truly)
}
