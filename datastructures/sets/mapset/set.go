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

type MapBasedSet[T any] struct {
	maps.Map[T, empty]
}

func NewHashSet[T comparable]() MapBasedSet[T] {
	return FromMap[T](hashmap.New[T, empty]())
}

func NewTreeSet[T cmp.Ordered]() MapBasedSet[T] {
	return FromMap[T](redblacktree.New[T, empty]())
}

func FromMap[T any](m maps.Map[T, empty]) MapBasedSet[T] {
	return MapBasedSet[T]{m}
}

func HashSetCollector[T comparable]() streams.Collector[T, MapBasedSet[T]] {
	return sets.Collector[T, MapBasedSet[T]]{NewHashSet[T]()}
}

func TreeSetCollector[T cmp.Ordered]() streams.Collector[T, MapBasedSet[T]] {
	return sets.Collector[T, MapBasedSet[T]]{NewTreeSet[T]()}
}

func (set MapBasedSet[T]) Add(value T) {
	if !set.Contains(value) {
		set.Map.Set(value, empty{})
	}
}

func (set MapBasedSet[T]) Clone() sets.Set[T] {
	return FromMap[T](set.Map.Clone())
}

func (set MapBasedSet[T]) Iterator() iterable.Iterator[T] {
	return set.ModifyingIterator()
}

func (set MapBasedSet[T]) ModifyingIterator() sets.Iterator[T] {
	return Iterator[T]{set.Map.ModifyingIterator()}
}

func (set MapBasedSet[T]) Stream() streams.Stream[T] {
	return streams.Map(set.Map.Stream(), func(pair misc.Pair[T, empty]) T {
		return pair.First
	})
}
