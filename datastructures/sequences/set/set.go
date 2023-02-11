package set

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type empty struct{}

type Set[T comparable] struct {
	m maps.Map[T, empty]
}

func NewFrom[T comparable](m maps.Map[T, empty]) Set[T] { return Set[T]{m} }

func New[T comparable]() Set[T] { return NewFrom[T](hashmap.New[T, empty]()) }

func (set Set[T]) Size() int { return set.m.Size() }

func (set Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.m.Set(value, empty{})
	}
}
func (set Set[T]) Remove(value T)        { set.m.Remove(value) }
func (set Set[T]) Contains(value T) bool { return set.m.Contains(value) }

func (set Set[T]) Clear()        { set.m.Clear() }
func (set Set[T]) Clone() Set[T] { return NewFrom[T](set.m.Clone()) }

func (set Set[T]) Iterator() iterable.Iterator[T] { return iterator[T]{set.m.Iterator()} }
func (set Set[T]) Stream() streams.Stream[T] {
	return streams.Map(set.m.Stream(), func(pair misc.Pair[T, empty]) T { return pair.First })
}
