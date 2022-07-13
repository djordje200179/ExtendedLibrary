package linkedlistset

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Set[T comparable] linkedlist.LinkedList[T]

func New[T comparable]() *Set[T] { return (*Set[T])(linkedlist.New[T]()) }

func Collector[T comparable]() streams.Collector[T, sets.Set[T]] { return sets.Collector[T](New[T]()) }

func (set *Set[T]) list() *linkedlist.LinkedList[T] { return (*linkedlist.LinkedList[T])(set) }

func (set *Set[T]) Size() int { return set.list().Size() }

func (set *Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.list().Append(value)
	}
}

func (set *Set[T]) Remove(value T) {
	for it := set.list().ModifyingIterator(); it.Valid(); it.Move() {
		if it.Get() == value {
			it.Remove()
			return
		}
	}
}

func (set *Set[T]) Contains(value T) bool {
	for it := set.list().Iterator(); it.Valid(); it.Move() {
		if it.Get() == value {
			return true
		}
	}

	return false
}

func (set *Set[T]) Clear() { set.list().Clear() }

func (set *Set[T]) Clone() sets.Set[T] {
	cloned := set.list().Clone().(*linkedlist.LinkedList[T])
	return (*Set[T])(cloned)
}

func (set *Set[T]) Iterator() datastructures.Iterator[T] { return set.list().Iterator() }
