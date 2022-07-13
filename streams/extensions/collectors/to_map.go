package collectors

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/bst"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/linkedlistmap"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type mapCollector[K comparable, V any] struct {
	m maps.Map[K, V]
}

func ToMap[K comparable, V any](empty maps.Map[K, V]) streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return mapCollector[K, V]{empty}
}

func (collector mapCollector[K, V]) Supply(value misc.Pair[K, V]) {
	collector.m.Set(value.First, value.Second)
}

func (collector mapCollector[K, V]) Finish() maps.Map[K, V] { return collector.m }

func ToBinarySearchTree[K comparable, V any](comparator functions.Comparator[K]) streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return ToMap[K, V](bst.New[K, V](comparator))
}

func ToHashMap[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return ToMap[K, V](hashmap.New[K, V]())
}

func ToLinkedListMap[K comparable, V any]() streams.Collector[misc.Pair[K, V], maps.Map[K, V]] {
	return ToMap[K, V](linkedlistmap.New[K, V]())
}
