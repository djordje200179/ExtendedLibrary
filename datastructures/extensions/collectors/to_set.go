package collectors

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/datastructures/sets/hashset"
	"github.com/djordje200179/extendedlibrary/datastructures/sets/linkedlistset"
	"github.com/djordje200179/extendedlibrary/streams"
)

type setCollector[T comparable] struct {
	set sets.Set[T]
}

func ToSet[T comparable](empty sets.Set[T]) streams.Collector[T, sets.Set[T]] {
	return setCollector[T]{empty}
}

func (collector setCollector[T]) Supply(value T) { collector.set.Add(value) }

func (collector setCollector[T]) Finish() sets.Set[T] { return collector.set }

func ToHashSet[T comparable]() streams.Collector[T, sets.Set[T]] {
	return ToSet[T](hashset.New[T]())
}

func ToLinkedListSet[T comparable]() streams.Collector[T, sets.Set[T]] {
	return ToSet[T](linkedlistset.New[T]())
}
