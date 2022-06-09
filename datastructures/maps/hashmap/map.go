package hashmap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Map[K comparable, V any] map[K]V

func New[K comparable, V any]() Map[K, V] {
	return make(map[K]V)
}

func (m Map[K, V]) Get(key K) V {
	return m[key]
}

func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m Map[K, V]) Remove(key K) {
	delete(m, key)
}

func (m Map[K, V]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

func (m Map[K, V]) Iterator() maps.Iterator[K, V] {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return &iterator[K, V]{
		m:     m,
		keys:  keys,
		index: 0,
	}
}

func (m Map[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	it := m.Iterator()

	return streams.Supply(func() misc.Pair[K, V] {
		defer it.Move()

		entry := it.Get()
		return misc.Pair[K, V]{entry.GetKey(), entry.GetValue()}
	})
}
