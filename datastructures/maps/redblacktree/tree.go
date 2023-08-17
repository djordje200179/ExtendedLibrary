package redblacktree

import (
	"cmp"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Tree[K, V any] struct {
	root  *Node[K, V]
	nodes int

	comparator comparison.Comparator[K]
}

func New[K, V any](comparator comparison.Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{
		comparator: comparator,
	}
}

func NewForOrderedKeys[K cmp.Ordered, V any]() *Tree[K, V] {
	return New[K, V](cmp.Compare[K])
}

func Collector[K, V any](comparator comparison.Comparator[K]) streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{New[K, V](comparator)}
}

func CollectorForOrderedKeys[K cmp.Ordered, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{NewForOrderedKeys[K, V]()}
}

func (tree *Tree[K, V]) Size() int {
	return tree.nodes
}

func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] {
	for curr := tree.root; curr != nil; {
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

func (tree *Tree[K, V]) Get(key K) V {
	return *tree.GetRef(key)
}

func (tree *Tree[K, V]) GetOrDefault(key K) V {
	return tree.GetOrElse(key, functions.Zero[V]())
}

func (tree *Tree[K, V]) GetOrElse(key K, value V) V {
	node := tree.GetNode(key)
	if node == nil {
		return value
	}

	return node.Value
}

func (tree *Tree[K, V]) TryGet(key K) (V, bool) {
	node := tree.GetNode(key)
	if node == nil {
		var zero V
		return zero, false
	}

	return node.Value, true
}

func (tree *Tree[K, V]) GetRef(key K) *V {
	node := tree.GetNode(key)
	if node == nil {
		maps.PanicOnMissingKey(key)
	}

	return &node.Value
}

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
		color:  red,
	}

	if tree.comparator(key, prev.key) == comparison.FirstSmaller {
		prev.left = node
	} else {
		prev.right = node
	}

	tree.nodes++

	tree.fixInsert(node)
}

func (tree *Tree[K, V]) Keys() []K {
	keys := make([]K, tree.nodes)

	i := 0
	for it := tree.Iterator(); it.Valid(); it.Move() {
		keys[i] = it.Get().First
		i++
	}

	return keys
}

func (tree *Tree[K, V]) Remove(key K) {
	node := tree.GetNode(key)
	if node == nil {
		return
	}

	tree.nodes--

	panic("implement me")
}

func (tree *Tree[K, V]) Contains(key K) bool {
	return tree.GetNode(key) != nil
}

func (tree *Tree[K, V]) Clear() {
	tree.root = nil
}

func (tree *Tree[K, V]) Swap(key1, key2 K) {
	node1, node2 := tree.GetNode(key1), tree.GetNode(key2)

	if node1 == nil {
		maps.PanicOnMissingKey(key1)
	}

	if node2 == nil {
		maps.PanicOnMissingKey(key2)
	}

	node1.Value, node2.Value = node2.Value, node1.Value
}

func (tree *Tree[K, V]) Clone() maps.Map[K, V] {
	//TODO implement me
	panic("implement me")
}

func (tree *Tree[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return tree.ModifyingIterator()
}

func (tree *Tree[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{tree, tree.root.Min()}
}

func (tree *Tree[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return iterable.IteratorStream(tree.Iterator())
}

func (tree *Tree[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](tree)
}

func (tree *Tree[K, V]) Root() *Node[K, V] {
	return tree.root
}
