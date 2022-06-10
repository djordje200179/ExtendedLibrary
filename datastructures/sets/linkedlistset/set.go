package linkedlistset

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Set[T comparable] *linkedlist.LinkedList[T]

func New[T comparable]() Set[T] {
	return linkedlist.New[T]()
}

func FromStream[T comparable](stream streams.Stream[T]) Set[T] {
	return linkedlist.FromStream(stream)
}

func (set Set[T]) getList() *linkedlist.LinkedList[T] {
	return set
}

func (set Set[T]) Add(value T) {
	if !set.Contains(value) {
		set.getList().Append(value)
	}
}

func (set Set[T]) Remove(value T) {
	for it := set.getList().Iterator(); it.IsValid(); it.Move() {
		if it.Get() == value {
			it.Remove()
			return
		}
	}
}

func (set Set[T]) Contains(value T) bool {
	for it := set.getList().Iterator(); it.IsValid(); it.Move() {
		if it.Get() == value {
			return true
		}
	}

	return false
}

func (set Set[T]) Iterator() datastructures.Iterator[T] {
	return set.getList().Iterator()
}

func (set Set[T]) Stream() streams.Stream[T] {
	return set.getList().Stream()
}
