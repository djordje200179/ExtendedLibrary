package hashmap

import "github.com/djordje200179/extendedlibrary/datastructures/maps"

type iterator[K comparable, V any] struct {
	m     HashMap[K, V]
	keys  []K
	index int
}

func (it *iterator[K, V]) Valid() bool { return it.index < len(it.keys) }
func (it *iterator[K, V]) Move()       { it.index++ }

func (it *iterator[K, V]) Get() maps.Entry[K, V] { return entry[K, V]{it.m, it.keys[it.index]} }

func (it *iterator[K, V]) Remove() { it.m.Remove(it.keys[it.index]) }
