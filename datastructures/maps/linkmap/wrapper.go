package linkmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iter"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/misc"
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

func (w *Wrapper[K, V]) moveToFront(node *Node[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		w.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		w.tail = node.prev
	}

	if w.head == nil {
		w.head = node
	} else {
		w.tail.next = node
		node.prev = w.tail
	}
	w.tail = node
}

// Size returns the number of entries in the map.
func (w *Wrapper[K, V]) Size() int {
	return w.m.Size()
}

// Contains returns true if the map contains the given key.
func (w *Wrapper[K, V]) Contains(key K) bool {
	return w.m.Contains(key)
}

// TryGet returns the value associated with the given key.
// If the key is not in the map, the zero value and false is returned.
// If the order is LRU, the entry is moved to the front.
func (w *Wrapper[K, V]) TryGet(key K) (V, bool) {
	node, ok := w.m.TryGet(key)
	if !ok {
		var zero V
		return zero, false
	}

	if w.order == LRU {
		w.moveToFront(node)
	}

	return node.Value, true
}

// Get returns the value associated with the given key.
// Panics if the key is not in the map.
// If the order is LRU, the entry is moved to the front.
func (w *Wrapper[K, V]) Get(key K) V {
	node := w.m.Get(key)

	if w.order == LRU {
		w.moveToFront(node)
	}

	return node.Value
}

// GetRef returns a reference to the value associated with the given key.
// Panics if the key is not in the map.
// If the order is LRU, the entry is moved to the front.
func (w *Wrapper[K, V]) GetRef(key K) *V {
	node := w.m.Get(key)

	if w.order == LRU {
		w.moveToFront(node)
	}

	return &node.Value
}

// Set sets the value associated with the given key.
// Entry is moved to the front of the map.
// If the map is full, the last entry is removed.
func (w *Wrapper[K, V]) Set(key K, value V) {
	node, ok := w.m.TryGet(key)
	if ok {
		node.Value = value

		if w.order == LRU {
			w.moveToFront(node)
		}

		return
	}

	if w.capacity > 0 && w.m.Size() == w.capacity {
		firstNode := w.head

		w.head = firstNode.next
		if w.head != nil {
			w.head.prev = nil
		} else {
			w.tail = nil
		}

		w.m.Remove(firstNode.key)
	}

	node = &Node[K, V]{key: key, Value: value}

	if w.head == nil {
		w.head = node
	} else {
		w.tail.next = node
		node.prev = w.tail
	}
	w.tail = node

	w.m.Set(key, node)
}

// Remove removes the entry with the given key.
// If the key is not in the map, nothing happens.
func (w *Wrapper[K, V]) Remove(key K) {
	node, ok := w.m.TryGet(key)
	if !ok {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		w.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		w.tail = node.prev
	}

	w.m.Remove(key)
}

// Clear removes all entries from the map.
func (w *Wrapper[K, V]) Clear() {
	w.m.Clear()

	w.head = nil
	w.tail = nil
}

// Clone returns a copy of the wrapper and
// of the underlying map.
func (w *Wrapper[K, V]) Clone() maps.Map[K, V] {
	clonedMap := w.m.Clone()

	var clonedHead, clonedTail *Node[K, V]
	for oldIt, newIt := w.m.Iterator(), clonedMap.MapIterator(); oldIt.Valid(); oldIt.Move() {
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
		capacity: w.capacity,
		order:    w.order,
	}
}

// Iterator returns an iter.Iterator over the entries in the map.
func (w *Wrapper[K, V]) Iterator() iter.Iterator[misc.Pair[K, V]] {
	return w.MapIterator()
}

// MapIterator returns a iterator over the entries in the map.
func (w *Wrapper[K, V]) MapIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{wrapper: w, curr: w.head}
}

// Stream2 streams over the entries in the Map.
func (w *Wrapper[K, V]) Stream2(yield func(K, V) bool) {
	for it := w.MapIterator(); it.Valid(); it.Move() {
		if !yield(it.Key(), it.Value()) {
			break
		}
	}
}

// Keys streams the keys of the maps.Map.
func (w *Wrapper[K, V]) Keys(yield func(K) bool) {
	for it := w.MapIterator(); it.Valid(); it.Move() {
		if !yield(it.Key()) {
			break
		}
	}
}

// Values streams the values of the maps.Map.
func (w *Wrapper[K, V]) Values(yield func(V) bool) {
	for it := w.MapIterator(); it.Valid(); it.Move() {
		if !yield(it.Value()) {
			break
		}
	}
}
