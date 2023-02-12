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

type mapBased[T comparable] struct {
	maps.Map[T, empty]
}

func NewFrom[T comparable](m maps.Map[T, empty]) sets.Set[T] {
	return mapBased[T]{m}
}

func NewHashSet[T comparable]() sets.Set[T] {
	return NewFrom[T](hashmap.New[T, empty]())
}

func Collector[T comparable]() streams.Collector[T, sets.Set[T]] {
	return sets.Collector[T](NewHashSet[T]())
}

func (set mapBased[T]) Add(value T) {
	if !set.Contains(value) {
		set.Map.Set(value, empty{})
	}
}

func (set mapBased[T]) Clone() sets.Set[T] {
	return NewFrom[T](set.Map.Clone())
}

func (set mapBased[T]) Iterator() iterable.Iterator[T] {
	return set.ModifyingIterator()
}

func (set mapBased[T]) ModifyingIterator() sets.Iterator[T] {
	return iterator[T]{
		Iterator: set.Map.ModifyingIterator(),
	}
}

func (set mapBased[T]) Stream() streams.Stream[T] {
	return streams.Map(set.Map.Stream(), func(pair misc.Pair[T, empty]) T {
		return pair.First
	})
}
