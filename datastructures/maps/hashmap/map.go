package hashmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type HashMap[K comparable, V any] map[K]V

func New[K comparable, V any]() HashMap[K, V]                 { return NewFromMap(make(map[K]V)) }
func NewFromMap[K comparable, V any](m map[K]V) HashMap[K, V] { return m }

func Collector[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return maps.Collector[K, V](New[K, V]())
}

func (hashmap HashMap[K, V]) Size() int { return len(hashmap) }

func (hashmap HashMap[K, V]) Get(key K) V        { return hashmap[key] }
func (hashmap HashMap[K, V]) GetRef(key K) *V    { panic("Not supported") }
func (hashmap HashMap[K, V]) Set(key K, value V) { hashmap[key] = value }

func (hashmap HashMap[K, V]) Remove(key K) { delete(hashmap, key) }

func (hashmap HashMap[K, V]) Contains(key K) bool {
	_, ok := hashmap[key]
	return ok
}

func (hashmap HashMap[K, V]) Clear() {
	for k := range hashmap {
		delete(hashmap, k)
	}
}

func (hashmap HashMap[K, V]) Clone() maps.Map[K, V] {
	cloned := New[K, V]()
	for k, v := range hashmap {
		cloned[k] = v
	}

	return cloned
}

func (hashmap HashMap[K, V]) Iterator() collections.Iterator[maps.Entry[K, V]] {
	return hashmap.ModifyingIterator()
}

func (hashmap HashMap[K, V]) ModifyingIterator() maps.Iterator[K, V] {
	keys := make([]K, len(hashmap))
	i := 0
	for k := range hashmap {
		keys[i] = k
		i++
	}

	return &iterator[K, V]{
		m:     hashmap,
		keys:  keys,
		index: 0,
	}
}

func (hashmap HashMap[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return streams.FromMap[K, V](hashmap)
}

func (hashmap HashMap[K, V]) RefStream() streams.Stream[misc.Pair[K, *V]] { panic("Not supported") }
