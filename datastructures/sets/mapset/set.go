package mapset

import (
	"cmp"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/redblacktree"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type empty struct{}

type Set[T any] struct {
	maps.Map[T, empty]
}

func NewHashSet[T comparable]() Set[T] {
	return FromMap[T](hashmap.New[T, empty]())
}

func NewTreeSet[T cmp.Ordered]() Set[T] {
	return FromMap[T](redblacktree.New[T, empty]())
}

func FromMap[T any](m maps.Map[T, empty]) Set[T] {
	return Set[T]{m}
}

func HashSetCollector[T comparable]() streams.Collector[T, Set[T]] {
	return sets.Collector[T, Set[T]]{NewHashSet[T]()}
}

func TreeSetCollector[T cmp.Ordered]() streams.Collector[T, Set[T]] {
	return sets.Collector[T, Set[T]]{NewTreeSet[T]()}
}

func (set Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.Map.Set(value, empty{})
	}
}

func (set Set[T]) Clone() sets.Set[T] {
	return FromMap[T](set.Map.Clone())
}

func (set Set[T]) Iterator() iterable.Iterator[T] {
	return set.ModifyingIterator()
}

func (set Set[T]) ModifyingIterator() sets.Iterator[T] {
	return Iterator[T]{set.Map.ModifyingIterator()}
}

func (set Set[T]) Stream() streams.Stream[T] {
	return streams.Map(set.Map.Stream(), func(pair misc.Pair[T, empty]) T {
		return pair.First
	})
}
