package linkedlistset

import (
	"github.com/djordje200179/GoExtendedLibrary/datastructures"
	"github.com/djordje200179/GoExtendedLibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/GoExtendedLibrary/streams"
)

type Set[T comparable] struct {
	list linkedlist.LinkedList[T]
}

func New[T comparable]() Set[T] {
	return Set[T]{linkedlist.New[T]()}
}

func (set *Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.list.Append(value)
	}
}

func (set *Set[T]) Remove(value T) {
	for it := set.list.Iterator(); it.IsValid(); it.Move() {
		if it.Get() == value {
			it.Remove()
			return
		}
	}
}

func (set *Set[T]) Contains(value T) bool {
	for it := set.list.Iterator(); it.IsValid(); it.Move() {
		if it.Get() == value {
			return true
		}
	}

	return false
}

func (set *Set[T]) Iterator() datastructures.Iterator[T] {
	return set.list.Iterator()
}

func (set *Set[T]) Stream() streams.Stream[T] {
	return set.list.Stream()
}