package rbt

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

// TODO: Implement Red-Black tree

type RedBlackTree[K comparable, V any] struct {
	root  *node[K, V]
	nodes int

	comparator functions.Comparator[K]
}

func New[K comparable, V any](comparator functions.Comparator[K]) *RedBlackTree[K, V] {
	tree := new(RedBlackTree[K, V])
	tree.root = nil
	tree.nodes = 0
	tree.comparator = comparator

	return tree
}

func Collector[K comparable, V any](comparator functions.Comparator[K]) streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V](New[K, V](comparator))
}

func (tree *RedBlackTree[K, V]) Size() int { return tree.nodes }

func (tree *RedBlackTree[K, V]) Get(key K) V {
	if ptr := tree.GetRef(key); ptr != nil {
		return *ptr
	} else {
		var empty V
		return empty
	}
}

func (tree *RedBlackTree[K, V]) GetRef(key K) *V {
	if node := tree.getNode(key); node != nil {
		return &node.value
	} else {
		return nil
	}
}

func (tree *RedBlackTree[K, V]) Set(key K, value V) {
	//TODO implement me
	panic("implement me")
}

func (tree *RedBlackTree[K, V]) Remove(key K) {
	//TODO implement me
	panic("implement me")
}

func (tree *RedBlackTree[K, V]) Contains(key K) bool { return tree.getNode(key) != nil }

func (tree *RedBlackTree[K, V]) Clear() {
	tree.root = nil
	tree.nodes = 0
}

func (tree *RedBlackTree[K, V]) Clone() maps.Map[K, V] {
	//TODO implement me
	panic("implement me")
}

func (tree *RedBlackTree[K, V]) Iterator() collections.Iterator[T] {
	return tree.ModifyingIterator()
}

func (tree *RedBlackTree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	//TODO implement me
	panic("implement me")
}

func (tree *RedBlackTree[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return streams.New(maps.ValuesStream[K, V](tree))
}

func (tree *RedBlackTree[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	return streams.New(maps.RefsStream[K, V](tree))
}
