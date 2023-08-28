package linkmap

// Node is a entry in the map.
// It is used to keep track of the order of the elements.
// It should not be created directly.
type Node[K, V any] struct {
	key   K
	Value V // The Value of the entry stored in the node.

	prev, next *Node[K, V]
}

// Key returns the key of the entry stored in the node.
func (node *Node[K, V]) Key() K {
	return node.key
}

// Prev returns the previous node in the map.
func (node *Node[K, V]) Prev() *Node[K, V] {
	return node.prev
}

// Next returns the next node in the map.
func (node *Node[K, V]) Next() *Node[K, V] {
	return node.next
}

// Clone returns a copy of the node.
// The clone has the same key and value as the node.
// The clone does not have any links to other nodes.
func (node *Node[K, V]) Clone() *Node[K, V] {
	return &Node[K, V]{
		key:   node.key,
		Value: node.Value,
	}
}
