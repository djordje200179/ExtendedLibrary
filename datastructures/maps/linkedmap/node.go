package linkedmap

type Node[K, V any] struct {
	key   K
	Value V

	prev, next *Node[K, V]
}

func (node *Node[K, V]) Key() K {
	return node.key
}

func (node *Node[K, V]) Prev() *Node[K, V] {
	return node.prev
}

func (node *Node[K, V]) Next() *Node[K, V] {
	return node.next
}

func (node *Node[K, V]) Clone() *Node[K, V] {
	return &Node[K, V]{
		key:   node.key,
		Value: node.Value,
	}
}
