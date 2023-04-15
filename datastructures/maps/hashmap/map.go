package hashmap

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	stdmaps "golang.org/x/exp/maps"
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

func (hashmap Map[K, V]) GetOrDefault(key K) V {
	return hashmap[key]
}

func (hashmap Map[K, V]) GetOrElse(key K, value V) V {
	if mapValue, ok := hashmap[key]; ok {
		return mapValue
	} else {
		return value
	}
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
	return stdmaps.Keys(hashmap)
}

func (hashmap Map[K, V]) Remove(key K) {
	delete(hashmap, key)
}

func (hashmap Map[K, V]) Contains(key K) bool {
	_, ok := hashmap[key]
	return ok
}

func (hashmap Map[K, V]) Clear() {
	mt, mv := mapTypeAndValue(hashmap)
	internalMapClear(mt, mv)
}

func (hashmap Map[K, V]) Swap(key1, key2 K) {
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
	return &iterator[K, V]{
		m:     hashmap,
		keys:  hashmap.Keys(),
		index: 0,
	}
}

func (hashmap Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return streams.FromMap[K, V](hashmap)
}

func (hashmap Map[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] {
	supplier := maps.RefsSupplier[K, V]{hashmap.ModifyingIterator()}
	return streams.Stream[misc.Pair[K, *V]]{supplier}
}

func (hashmap Map[K, V]) Map() map[K]V {
	return hashmap
}
