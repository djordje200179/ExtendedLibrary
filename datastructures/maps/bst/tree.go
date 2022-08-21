package bst

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/linears/queue"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
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
	return maps.Collector[K, V](New[K, V](comparator))
}

func (tree *BinarySearchTree[K, V]) Size() int { return tree.nodes }

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

func (tree *BinarySearchTree[K, V]) Remove(key K) {
	if node := tree.getNode(key); node != nil {
		tree.removeNode(node)
	}
}

func (tree *BinarySearchTree[K, V]) Contains(key K) bool { return tree.getNode(key) != nil }

func (tree *BinarySearchTree[K, V]) Clear() {
	tree.root = nil
	tree.nodes = 0
}

func (tree *BinarySearchTree[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V](tree.comparator)
	cloned.nodes = tree.nodes

	if tree.root == nil {
		return cloned
	}

	cloned.root = tree.root.Clone()

	nodesInOriginal := queue.New[*node[K, V]]()
	nodesInOriginal.Push(tree.root)

	nodesInCloned := queue.New[*node[K, V]]()
	nodesInCloned.Push(cloned.root)

	for !nodesInOriginal.Empty() {
		nodeInOriginal := nodesInOriginal.Pop()
		nodeInCloned := nodesInCloned.Pop()

		if left := nodeInOriginal.left; left != nil {
			nodesInOriginal.Push(left)
			nodeInCloned.left = left.Clone()
		}

		if right := nodeInOriginal.right; right != nil {
			nodesInOriginal.Push(right)
			nodeInCloned.right = right.Clone()
		}
	}

	return cloned
}

func (tree *BinarySearchTree[K, V]) Iterator() collections.Iterator[maps.Entry[K, V]] {
	return tree.ModifyingIterator()
}

func (tree *BinarySearchTree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &iterator[K, V]{tree, tree.root.min()}
}

func (tree *BinarySearchTree[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return streams.New(maps.ValuesSupplier[K, V](tree))
}

func (tree *BinarySearchTree[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	return streams.New(maps.RefsSupplier[K, V](tree))
}
