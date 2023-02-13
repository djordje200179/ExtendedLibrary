package bst

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/collectionsequences"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type BinarySearchTree[K comparable, V any] struct {
	root  *node[K, V]
	nodes int

	comparator functions.Comparator[K]
}

func New[K comparable, V any](comparator functions.Comparator[K]) *BinarySearchTree[K, V] {
	tree := new(BinarySearchTree[K, V])
	tree.root = nil
	tree.nodes = 0
	tree.comparator = comparator

	return tree
}

func Collector[K comparable, V any](comparator functions.Comparator[K]) streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{
		Map: New[K, V](comparator),
	}
}

func (tree *BinarySearchTree[K, V]) Size() int {
	return tree.nodes
}

func (tree *BinarySearchTree[K, V]) Get(key K) V {
	if ptr := tree.GetRef(key); ptr != nil {
		return *ptr
	} else {
		var empty V
		return empty
	}
}

func (tree *BinarySearchTree[K, V]) GetRef(key K) *V {
	if node := tree.getNode(key); node != nil {
		return &node.value
	} else {
		return nil
	}
}

func (tree *BinarySearchTree[K, V]) Set(key K, value V) {
	if tree.root == nil {
		tree.root = &node[K, V]{
			key:   key,
			value: value,
		}
		tree.nodes++

		return
	}

	prev := (*node[K, V])(nil)
	for curr := tree.root; curr != nil; {
		if key == curr.key {
			curr.value = value
			return
		}

		prev = curr
		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.left
		case comparison.FirstBigger:
			curr = curr.right
		case comparison.Equal:
			curr.value = value
			return
		}
	}

	node := &node[K, V]{
		key:    key,
		value:  value,
		parent: prev,
	}

	if tree.comparator(key, prev.key) == comparison.FirstSmaller {
		prev.left = node
	} else {
		prev.right = node
	}

	tree.nodes++
}

func (tree *BinarySearchTree[K, V]) Keys() []K {
	keys := make([]K, tree.nodes)

	i := 0
	for it := tree.Iterator(); it.Valid(); it.Move() {
		keys[i] = it.Get().Key()
		i++
	}

	return keys
}

func (tree *BinarySearchTree[K, V]) Remove(key K) {
	if node := tree.getNode(key); node != nil {
		tree.removeNode(node)
	}
}

func (tree *BinarySearchTree[K, V]) Contains(key K) bool {
	return tree.getNode(key) != nil
}

func (tree *BinarySearchTree[K, V]) Clear() {
	tree.root = nil
	tree.nodes = 0
}

func (tree *BinarySearchTree[K, V]) Swap(key1, key2 K) {
	node1, node2 := tree.getNode(key1), tree.getNode(key2)

	node1.value, node2.value = node2.value, node1.value
}

func (tree *BinarySearchTree[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V](tree.comparator)
	cloned.nodes = tree.nodes

	if tree.root == nil {
		return cloned
	}

	cloned.root = tree.root.Clone()

	nodesInOriginal := collectionsequences.NewQueue[*node[K, V]]()
	nodesInOriginal.PushBack(tree.root)

	nodesInCloned := collectionsequences.NewQueue[*node[K, V]]()
	nodesInCloned.PushBack(cloned.root)

	for !nodesInOriginal.Empty() {
		nodeInOriginal := nodesInOriginal.PopFront()
		nodeInCloned := nodesInCloned.PopFront()

		if left := nodeInOriginal.left; left != nil {
			nodesInOriginal.PushBack(left)
			nodeInCloned.left = left.Clone()
		}

		if right := nodeInOriginal.right; right != nil {
			nodesInOriginal.PushBack(right)
			nodeInCloned.right = right.Clone()
		}
	}

	return cloned
}

func (tree *BinarySearchTree[K, V]) Iterator() iterable.Iterator[K] {
	return tree.ModifyingIterator()
}

func (tree *BinarySearchTree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &iterator[K, V]{tree, tree.root.min()}
}

func (tree *BinarySearchTree[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return maps.ValuesStream[K, V](tree)
}

func (tree *BinarySearchTree[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](tree)
}
