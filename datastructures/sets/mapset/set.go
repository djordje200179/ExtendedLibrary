package mapset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type empty struct{}

type mapBasedSet[T comparable] struct {
	maps.Map[T, empty]
}

func NewHashSet[T comparable]() sets.Set[T] {
	return From[T](hashmap.New[T, empty]())
}

func From[T comparable](m maps.Map[T, empty]) sets.Set[T] {
	return mapBasedSet[T]{m}
}

func HashSetCollector[T comparable]() streams.Collector[T, sets.Set[T]] {
	return sets.Collector[T]{NewHashSet[T]()}
}

func (set mapBasedSet[T]) Add(value T) {
	if !set.Contains(value) {
		set.Map.Set(value, empty{})
	}
}

func (set mapBasedSet[T]) Clone() sets.Set[T] {
	return From[T](set.Map.Clone())
}

func (set mapBasedSet[T]) Iterator() iterable.Iterator[T] {
	return set.ModifyingIterator()
}

func (set mapBasedSet[T]) ModifyingIterator() sets.Iterator[T] {
	return iterator[T]{set.Map.ModifyingIterator()}
}

func (set mapBasedSet[T]) Stream() streams.Stream[T] {
	return streams.Map(set.Map.Stream(), func(pair misc.Pair[T, empty]) T {
		return pair.First
	})
}
