package linkedmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Wrapper[K, V any] struct {
	m maps.Map[K, *Node[K, V]]

	head, tail *Node[K, V]

	capacity int
}

func From[K, V any](m maps.Map[K, *Node[K, V]], capacity int) *Wrapper[K, V] {
	wrapper := &Wrapper[K, V]{
		m:        m,
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

func NewHashmap[K comparable, V any]() *Wrapper[K, V] {
	return NewFIFOHashmap[K, V](0)
}

func NewFIFOHashmap[K comparable, V any](capacity int) *Wrapper[K, V] {
	m := hashmap.NewWithCapacity[K, *Node[K, V]](capacity)
	wrapper := &Wrapper[K, V]{
		m:        m,
		capacity: capacity,
	}

	return wrapper
}

func (wrapper *Wrapper[K, V]) Size() int {
	return wrapper.m.Size()
}

func (wrapper *Wrapper[K, V]) Contains(key K) bool {
	return wrapper.m.Contains(key)
}

func (wrapper *Wrapper[K, V]) TryGet(key K) (V, bool) {
	node, ok := wrapper.m.TryGet(key)
	if !ok {
		return functions.Zero[V](), false
	}

	return node.Value, true
}

func (wrapper *Wrapper[K, V]) Get(key K) V {
	return wrapper.m.Get(key).Value
}

func (wrapper *Wrapper[K, V]) GetRef(key K) *V {
	return &wrapper.m.Get(key).Value
}

func (wrapper *Wrapper[K, V]) Set(key K, value V) {
	node, ok := wrapper.m.TryGet(key)
	if ok {
		node.Value = value
		return
	}

	if wrapper.capacity > 0 && wrapper.m.Size() > wrapper.capacity {
		firstNode := wrapper.head

		wrapper.head = firstNode.next
		if wrapper.head != nil {
			wrapper.head.prev = nil
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

func (wrapper *Wrapper[K, V]) Keys() []K {
	return wrapper.m.Keys()
}

func (wrapper *Wrapper[K, V]) Clear() {
	wrapper.m.Clear()

	wrapper.head = nil
	wrapper.tail = nil
}

func (wrapper *Wrapper[K, V]) Clone() maps.Map[K, V] {
	clonedMap := wrapper.m.Clone()

	var clonedHead, clonedTail *Node[K, V]
	for oldIt, newIt := wrapper.m.Iterator(), clonedMap.ModifyingIterator(); oldIt.Valid(); oldIt.Move() {
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
		m:        clonedMap,
		head:     clonedHead,
		tail:     clonedTail,
		capacity: wrapper.capacity,
	}
}

func (wrapper *Wrapper[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return wrapper.ModifyingIterator()
}

func (wrapper *Wrapper[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{wrapper: wrapper, curr: wrapper.head}
}

func (wrapper *Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return iterable.IteratorStream(wrapper.Iterator())
}

func (wrapper *Wrapper[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](wrapper)
}
