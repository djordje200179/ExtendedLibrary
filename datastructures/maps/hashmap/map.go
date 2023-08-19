package hashmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"unsafe"
)

type Map[K comparable, V any] map[K]V

func New[K comparable, V any]() Map[K, V] {
	return NewWithCapacity[K, V](0)
}

func NewWithCapacity[K comparable, V any](capacity int) Map[K, V] {
	return FromMap(make(map[K]V, capacity))
}

func FromMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V]{New[K, V]()}
}

func (hashmap Map[K, V]) Size() int {
	return len(hashmap)
}

func (hashmap Map[K, V]) Get(key K) V {
	value, ok := hashmap[key]
	if !ok {
		maps.PanicOnMissingKey(key)
	}

	return value
}

func (hashmap Map[K, V]) TryGet(key K) (V, bool) {
	value, ok := hashmap[key]
	return value, ok
}

func (hashmap Map[K, V]) GetRef(key K) *V {
	mt, mv := mapTypeAndValue(hashmap)
	ptr, ok := internalMapGet(mt, mv, unsafe.Pointer(&key))

	if !ok {
		maps.PanicOnMissingKey(key)
	}

	return (*V)(ptr)
}

func (hashmap Map[K, V]) Set(key K, value V) {
	hashmap[key] = value
}

func (hashmap Map[K, V]) Keys() []K {
	keys := make([]K, len(hashmap))

	i := 0
	for key := range hashmap {
		keys[i] = key
		i++
	}

	return keys
}

func (hashmap Map[K, V]) Remove(key K) {
	delete(hashmap, key)
}

func (hashmap Map[K, V]) Contains(key K) bool {
	_, ok := hashmap[key]
	return ok
}

func (hashmap Map[K, V]) Clear() {
	clear(hashmap)
}

func (hashmap Map[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V]()
	for k, v := range hashmap {
		cloned[k] = v
	}

	return cloned
}

func (hashmap Map[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return hashmap.ModifyingIterator()
}

func (hashmap Map[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	return &Iterator[K, V]{
		m:     hashmap,
		keys:  hashmap.Keys(),
		index: 0,
	}
}

func (hashmap Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	supplier := suppliers.Map(hashmap)
	return streams.New(supplier)
}

func (hashmap Map[K, V]) RefsStream() streams.Stream[misc.Pair[K, *V]] {
	return maps.RefsStream[K, V](hashmap)
}

func (hashmap Map[K, V]) Map() map[K]V {
	return hashmap
}
