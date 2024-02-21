package rbt

import (
	"cmp"
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/seqs/colseq"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

// Tree is a red-black tree implementation of a map.
// The zero value is ready to use. Do not copy a non-zero Tree.
type Tree[K, V any] struct {
	root  *Node[K, V]
	nodes int

	comparator comparison.Comparator[K]
}

// NewWithComparator creates an empty Tree with the specified comparator.
func NewWithComparator[K, V any](comparator comparison.Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{
		comparator: comparator,
	}
}

// New creates an empty Tree with the default comparator for ordered keys.
func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return NewWithComparator[K, V](cmp.Compare[K])
}

// NewWithComparatorFromIterable creates a Tree with the specified comparator from the specified iter.Iterable.
func NewWithComparatorFromIterable[K, V any](comparator comparison.Comparator[K], iterable iter.Iterable[misc.Pair[K, V]]) *Tree[K, V] {
	tree := NewWithComparator[K, V](comparator)

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		entry := it.Get()
		tree.Set(entry.First, entry.Second)
	}

	return tree
}

// NewFromIterable creates a Tree from the specified iter.Iterable.
func NewFromIterable[K cmp.Ordered, V any](iterable iter.Iterable[misc.Pair[K, V]]) *Tree[K, V] {
	tree := New[K, V]()

	for it := iterable.Iterator(); it.Valid(); it.Move() {
		entry := it.Get()
		tree.Set(entry.First, entry.Second)
	}

	return tree
}

// Size returns the number of entries in the tree.
func (tree *Tree[K, V]) Size() int {
	return tree.nodes
}

// GetNode returns the node associated with the specified key.
// Returns nil if the key is not present.
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] {
	for curr := tree.root; curr != nil; {
		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.leftChild
		case comparison.FirstBigger:
			curr = curr.rightChild
		case comparison.Equal:
			return curr
		}
	}

	return nil
}

// Contains returns true if the tree contains the specified key.
func (tree *Tree[K, V]) Contains(key K) bool {
	return tree.GetNode(key) != nil
}

// TryGet returns the value associated with the specified key,
// or zero value and false if the key is not present.
func (tree *Tree[K, V]) TryGet(key K) (V, bool) {
	node := tree.GetNode(key)
	if node == nil {
		var zero V
		return zero, false
	}

	return node.Value, true
}

// Get returns the value associated with the specified key.
// Panics if the key is not present.
func (tree *Tree[K, V]) Get(key K) V {
	return *tree.GetRef(key)
}

// GetRef returns a reference to the value associated with the specified key.
// Panics if the key is not present.
func (tree *Tree[K, V]) GetRef(key K) *V {
	node := tree.GetNode(key)
	if node == nil {
		panic(maps.ErrMissingKey[K]{Key: key})
	}

	return &node.Value
}

// Set sets the value associated with the specified key.
func (tree *Tree[K, V]) Set(key K, value V) {
	if tree.root == nil {
		tree.root = &Node[K, V]{
			key:   key,
			Value: value,
			color: black,
		}

		tree.nodes++

		return
	}

	var prev *Node[K, V]
	for curr := tree.root; curr != nil; {
		prev = curr
		switch tree.comparator(key, curr.key) {
		case comparison.FirstSmaller:
			curr = curr.leftChild
		case comparison.FirstBigger:
			curr = curr.rightChild
		case comparison.Equal:
			curr.Value = value
			return
		}
	}

	node := &Node[K, V]{
		key:    key,
		Value:  value,
		parent: prev,
		color:  red,
	}

	if tree.comparator(key, prev.key) == comparison.FirstSmaller {
		prev.leftChild = node
	} else {
		prev.rightChild = node
	}

	tree.nodes++

	tree.fixInsert(node)
}

func (tree *Tree[K, V]) removeNode(node *Node[K, V]) {
	if node.leftChild != nil && node.rightChild != nil {
		next := node.Next()

		next.key = node.key
		next.Value = node.Value

		tree.removeNode(next)

		return
	}

	tree.nodes--

	var child *Node[K, V]
	if node.leftChild != nil {
		child = node.leftChild
	} else {
		child = node.rightChild
	}

	var locationInParent **Node[K, V]
	if node.parent == nil {
		locationInParent = &tree.root
	} else if node.parent.leftChild == node {
		locationInParent = &node.parent.leftChild
	} else {
		locationInParent = &node.parent.rightChild
	}

	if child != nil {
		child.parent = node.parent
		*locationInParent = child

		if node.color == black {
			tree.fixRemove(child)
		}
	} else if node.parent == nil {
		*locationInParent = nil
	} else {
		if node.color == black {
			tree.fixRemove(node)
		}

		if node.parent != nil {
			if node.parent.leftChild == node {
				node.parent.leftChild = nil
			} else {
				node.parent.rightChild = nil
			}

			node.parent = nil
		}
	}
}

// Remove removes the entry with the specified key.
// Does nothing if the key is not present.
func (tree *Tree[K, V]) Remove(key K) {
	node := tree.GetNode(key)
	if node != nil {
		tree.removeNode(node)
	}
}

// Clear removes all entries from the tree.
func (tree *Tree[K, V]) Clear() {
	tree.root = nil
}

// Clone returns a shallow copy of the tree.
func (tree *Tree[K, V]) Clone() maps.Map[K, V] {
	newTree := &Tree[K, V]{
		nodes:      tree.nodes,
		comparator: tree.comparator,
	}

	if tree.root == nil {
		return newTree
	}

	newTree.root = tree.root.Clone()

	nodesInOriginal := colseq.NewArrayDeque[*Node[K, V]]()
	nodesInOriginal.PushBack(tree.root)

	nodesInCloned := colseq.NewArrayDeque[*Node[K, V]]()
	nodesInCloned.PushBack(newTree.root)

	for !nodesInOriginal.Empty() {
		nodeInOriginal := nodesInOriginal.PopFront()
		nodeInCloned := nodesInCloned.PopFront()

		if leftNode := nodeInOriginal.leftChild; leftNode != nil {
			nodesInOriginal.PushBack(leftNode)

			newLeftNode := leftNode.Clone()
			nodeInCloned.leftChild = newLeftNode
			newLeftNode.parent = nodeInCloned
		}

		if rightNode := nodeInOriginal.rightChild; rightNode != nil {
			nodesInOriginal.PushBack(rightNode)

			newRightNode := rightNode.Clone()
			nodeInCloned.rightChild = newRightNode
			newRightNode.parent = nodeInCloned
		}
	}

	return newTree
}

// Iterator returns an iter.Iterator over the tree.
func (tree *Tree[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return tree.MapIterator()
}

// MapIterator returns an iterator over the tree.
func (tree *Tree[K, V]) MapIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{tree, tree.root.Min()}
}

// Stream2 streams over the entries in the Tree.
func (tree *Tree[K, V]) Stream2(yield func(K, V) bool) {
	for it := tree.MapIterator(); it.Valid(); it.Move() {
		if !yield(it.Key(), it.Value()) {
			return
		}
	}
}

// RefsStream2 streams over the keys and references to the values in the Map.
func (tree *Tree[K, V]) RefsStream2(yield func(K, *V) bool) {
	for it := tree.MapIterator(); it.Valid(); it.Move() {
		if !yield(it.Key(), it.ValueRef()) {
			return
		}
	}
}

// Root returns the root node of the tree.
func (tree *Tree[K, V]) Root() *Node[K, V] {
	return tree.root
}
