package binarysearchtree

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/collectionsequences"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type BinarySearchTree[K comparable, V any] struct {
	root  *Node[K, V]
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

func (tree *BinarySearchTree[K, V]) GetNode(key K) *Node[K, V] {
	for curr := tree.root; curr != nil; {
		if key == curr.key {
			return curr
		}

		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.left
		case comparison.FirstBigger:
			curr = curr.right
		case comparison.Equal:
			return curr
		}
	}

	return nil
}

func (tree *BinarySearchTree[K, V]) GetRef(key K) *V {
	node := tree.GetNode(key)
	if node == nil {
		panic(fmt.Sprintf("Key %v not found", key))
	}

	return &node.Value
}

func (tree *BinarySearchTree[K, V]) Get(key K) V {
	return *tree.GetRef(key)
}

func (tree *BinarySearchTree[K, V]) Set(key K, value V) {
	if tree.root == nil {
		tree.root = &Node[K, V]{
			key:   key,
			Value: value,
		}

		tree.nodes++

		return
	}

	var prev *Node[K, V]
	for curr := tree.root; curr != nil; {
		if key == curr.key {
			curr.Value = value
			return
		}

		prev = curr
		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.left
		case comparison.FirstBigger:
			curr = curr.right
		case comparison.Equal:
			curr.Value = value
			return
		}
	}

	node := &Node[K, V]{
		key:    key,
		Value:  value,
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
		keys[i] = it.Get().First
		i++
	}

	return keys
}

func (tree *BinarySearchTree[K, V]) removeNode(node *Node[K, V]) {
	locationInParent := node.locationInParent()
	if locationInParent == nil {
		locationInParent = &tree.root
	}

	if node.left == nil && node.right == nil {
		*locationInParent = nil
	} else if node.left == nil {
		*locationInParent = node.right
	} else if node.right == nil {
		*locationInParent = node.left
	} else {
		next := node.Next()

		node.key, next.key = next.key, node.key
		node.Value, next.Value = next.Value, node.Value

		tree.removeNode(next)
	}

	tree.nodes--
}

func (tree *BinarySearchTree[K, V]) Remove(key K) {
	node := tree.GetNode(key)
	if node == nil {
		panic(fmt.Sprintf("Key %v not found", key))
	}

	tree.removeNode(node)
}

func (tree *BinarySearchTree[K, V]) Contains(key K) bool {
	return tree.GetNode(key) != nil
}

func (tree *BinarySearchTree[K, V]) Clear() {
	tree.root = nil
	tree.nodes = 0
}

func (tree *BinarySearchTree[K, V]) Swap(key1, key2 K) {
	node1, node2 := tree.GetNode(key1), tree.GetNode(key2)

	node1.Value, node2.Value = node2.Value, node1.Value
}

func (tree *BinarySearchTree[K, V]) Clone() maps.Map[K, V] {
	cloned := &BinarySearchTree[K, V]{
		nodes:      tree.nodes,
		comparator: tree.comparator,
	}

	if tree.root == nil {
		return cloned
	}

	cloned.root = tree.root.Clone()

	nodesInOriginal := collectionsequences.NewDeque[*Node[K, V]]()
	nodesInOriginal.PushBack(tree.root)

	nodesInCloned := collectionsequences.NewDeque[*Node[K, V]]()
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

func (tree *BinarySearchTree[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return tree.ModifyingIterator()
}

func (tree *BinarySearchTree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &iterator[K, V]{tree, tree.root.Min()}
}

func (tree *BinarySearchTree[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	supplier := iterable.IteratorSupplier[misc.Pair[K, V]]{
		Iterator: tree.Iterator(),
	}

	return streams.Stream[misc.Pair[K, V]]{
		Supplier: supplier,
	}
}

func (tree *BinarySearchTree[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	supplier := maps.RefsSupplier[K, V]{
		Iterator: tree.ModifyingIterator(),
	}

	return streams.Stream[misc.Pair[K, *V]]{
		Supplier: supplier,
	}
}

func (tree *BinarySearchTree[K, V]) Root() *Node[K, V] {
	return tree.root
}
