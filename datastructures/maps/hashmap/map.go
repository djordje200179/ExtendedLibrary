package hashmap

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type HashMap[K comparable, V any] map[K]V

func New[K comparable, V any]() HashMap[K, V] {
	return NewWithCapacity[K, V](0)
}

func NewWithCapacity[K comparable, V any](capacity int) HashMap[K, V] {
	return From(make(map[K]V, capacity))
}

func From[K comparable, V any](m map[K]V) HashMap[K, V] {
	return m
}

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{New[K, V]()}
}

func (hashmap HashMap[K, V]) Size() int {
	return len(hashmap)
}

func (hashmap HashMap[K, V]) Get(key K) V {
	value, ok := hashmap[key]
	if !ok {
		panic(fmt.Sprintf("Key %v not found", key))
	}

	return value
}

func (hashmap HashMap[K, V]) GetRef(key K) *V {
	panic("Getting reference to value from hash map is not supported")
}

func (hashmap HashMap[K, V]) Set(key K, value V) {
	hashmap[key] = value
}

func (hashmap HashMap[K, V]) Keys() []K {
	keys := make([]K, len(hashmap))

	i := 0
	for k := range hashmap {
		keys[i] = k
		i++
	}

	return keys
}

func (hashmap HashMap[K, V]) Remove(key K) {
	delete(hashmap, key)
}

func (hashmap HashMap[K, V]) Contains(key K) bool {
	_, ok := hashmap[key]
	return ok
}

func (hashmap HashMap[K, V]) Clear() {
	for k := range hashmap {
		delete(hashmap, k)
	}
}

func (hashmap HashMap[K, V]) Swap(key1, key2 K) {
	value1, ok1 := hashmap[key1]
	if !ok1 {
		panic(fmt.Sprintf("Key %v not found", key1))
	}

	value2, ok2 := hashmap[key2]
	if !ok2 {
		panic(fmt.Sprintf("Key %v not found", key2))
	}

	hashmap[key1], hashmap[key2] = value2, value1
}

func (hashmap HashMap[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V]()
	for k, v := range hashmap {
		cloned[k] = v
	}

	return cloned
}

func (hashmap HashMap[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return hashmap.ModifyingIterator()
}

func (hashmap HashMap[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &iterator[K, V]{
		m:     hashmap,
		keys:  hashmap.Keys(),
		index: 0,
	}
}

func (hashmap HashMap[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return streams.FromMap[K, V](hashmap)
}

func (hashmap HashMap[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	panic("Getting reference to value from hash map is not supported")
}

func (hashmap HashMap[K, V]) Map() map[K]V {
	return hashmap
}
