package hashset

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
)

type Set[T comparable] struct {
	m hashmap.Map[T, bool]
}

func New[T comparable]() Set[T] {
	return Set[T]{hashmap.New[T, bool]()}
}

func (set Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.m.Set(value, true)
	}
}

func (set Set[T]) Remove(value T) {
	delete(set.m, value)
}

func (set Set[T]) Contains(value T) bool {
	return set.m.Contains(value)
}

func (set Set[T]) Iterator() datastructures.Iterator[T] {
	return iterator[T]{set.m.Iterator()}
}
