package linkmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

// Order represents the order of the nodes in the map.
type Order bool

const (
	FIFO Order = false // FIFO represents the first-in-first-out order.
	LRU        = true  // LRU represents the least-recently-used order.
)

// Wrapper is a map that keeps track of the order of the nodes.
type Wrapper[K, V any] struct {
	m maps.Map[K, *Node[K, V]]

	head, tail *Node[K, V]
	order      Order

	capacity int
}

// From returns a new Wrapper that wraps the given map.
func From[K, V any](m maps.Map[K, *Node[K, V]], capacity int, order Order) *Wrapper[K, V] {
	wrapper := &Wrapper[K, V]{
		m: m,

		order:    order,
		capacity: capacity,
	}

	for it := m.Iterator(); it.Valid(); it.Move() {
		node := it.Get().Second
		if wrapper.head == nil {
			wrapper.head = node
		} else {
			wrapper.tail = node
		}
	}

	return wrapper
}

// NewHashmap returns a new Wrapper around a empty hashmap.
func NewHashmap[K comparable, V any](capacity int, order Order) *Wrapper[K, V] {
	m := hashmap.NewWithCapacity[K, *Node[K, V]](capacity)
	wrapper := &Wrapper[K, V]{
		m: m,

		order:    order,
		capacity: capacity,
	}

	return wrapper
}

func (wrapper *Wrapper[K, V]) moveToFront(node *Node[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		wrapper.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		wrapper.tail = node.prev
	}

	if wrapper.head == nil {
		wrapper.head = node
	} else {
		wrapper.tail.next = node
		node.prev = wrapper.tail
	}
	wrapper.tail = node
}

// Size returns the number of entries in the map.
func (wrapper *Wrapper[K, V]) Size() int {
	return wrapper.m.Size()
}

// Contains returns true if the map contains the given key.
func (wrapper *Wrapper[K, V]) Contains(key K) bool {
	return wrapper.m.Contains(key)
}

// TryGet returns the value associated with the given key.
// If the key is not in the map, the zero value and false is returned.
// If the order is LRU, the entry is moved to the front.
func (wrapper *Wrapper[K, V]) TryGet(key K) (V, bool) {
	node, ok := wrapper.m.TryGet(key)
	if !ok {
		return functions.Zero[V](), false
	}

	if wrapper.order == LRU {
		wrapper.moveToFront(node)
	}

	return node.Value, true
}

// Get returns the value associated with the given key.
// Panics if the key is not in the map.
// If the order is LRU, the entry is moved to the front.
func (wrapper *Wrapper[K, V]) Get(key K) V {
	node := wrapper.m.Get(key)

	if wrapper.order == LRU {
		wrapper.moveToFront(node)
	}

	return node.Value
}

// GetRef returns a reference to the value associated with the given key.
// Panics if the key is not in the map.
// If the order is LRU, the entry is moved to the front.
func (wrapper *Wrapper[K, V]) GetRef(key K) *V {
	node := wrapper.m.Get(key)

	if wrapper.order == LRU {
		wrapper.moveToFront(node)
	}

	return &node.Value
}

// Set sets the value associated with the given key.
// Entry is moved to the front of the map.
// If the map is full, the last entry is removed.
func (wrapper *Wrapper[K, V]) Set(key K, value V) {
	node, ok := wrapper.m.TryGet(key)
	if ok {
		node.Value = value

		if wrapper.order == LRU {
			wrapper.moveToFront(node)
		}

		return
	}

	if wrapper.capacity > 0 && wrapper.m.Size() == wrapper.capacity {
		firstNode := wrapper.head

		wrapper.head = firstNode.next
		if wrapper.head != nil {
			wrapper.head.prev = nil
		} else {
			wrapper.tail = nil
		}

		wrapper.m.Remove(firstNode.key)
	}

	node = &Node[K, V]{key: key, Value: value}

	if wrapper.head == nil {
		wrapper.head = node
	} else {
		wrapper.tail.next = node
		node.prev = wrapper.tail
	}
	wrapper.tail = node

	wrapper.m.Set(key, node)
}

// Remove removes the entry with the given key.
// If the key is not in the map, nothing happens.
func (wrapper *Wrapper[K, V]) Remove(key K) {
	node, ok := wrapper.m.TryGet(key)
	if !ok {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		wrapper.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		wrapper.tail = node.prev
	}

	wrapper.m.Remove(key)
}

// Keys returns a slice of all keys in the map.
func (wrapper *Wrapper[K, V]) Keys() []K {
	return wrapper.m.Keys()
}

// Clear removes all entries from the map.
func (wrapper *Wrapper[K, V]) Clear() {
	wrapper.m.Clear()

	wrapper.head = nil
	wrapper.tail = nil
}

// Clone returns a copy of the wrapper and
// of the underlying map.
func (wrapper *Wrapper[K, V]) Clone() maps.Map[K, V] {
	clonedMap := wrapper.m.Clone()

	var clonedHead, clonedTail *Node[K, V]
	for oldIt, newIt := wrapper.m.Iterator(), clonedMap.MapIterator(); oldIt.Valid(); oldIt.Move() {
		oldNode := oldIt.Get().Second

		newNode := oldNode.Clone()
		if clonedHead == nil {
			clonedHead = newNode
		} else {
			clonedTail.next = newNode
			newNode.prev = clonedTail
		}
		clonedTail = newNode

		newIt.SetValue(newNode)

		newIt.Move()
	}

	return &Wrapper[K, V]{
		m: clonedMap,

		head:     clonedHead,
		tail:     clonedTail,
		capacity: wrapper.capacity,
		order:    wrapper.order,
	}
}

// Iterator returns an iter.Iterator over the entries in the map.
func (wrapper *Wrapper[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return wrapper.MapIterator()
}

// MapIterator returns a iterator over the entries in the map.
func (wrapper *Wrapper[K, V]) MapIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{wrapper: wrapper, curr: wrapper.head}
}

// Stream returns a streams.Stream over the entries in the map.
func (wrapper *Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return iter.IteratorStream(wrapper.Iterator())
}

// RefsStream returns a streams.Stream over the entries in the map.
func (wrapper *Wrapper[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](wrapper)
}
