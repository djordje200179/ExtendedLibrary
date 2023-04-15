package readonlymap

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Wrapper[K comparable, V any] struct {
	m maps.Map[K, V]
}

func From[K comparable, V any](m maps.Map[K, V]) Wrapper[K, V] {
	return Wrapper[K, V]{m}
}

func (wrapper Wrapper[K, V]) Size() int {
	return wrapper.m.Size()
}

func (wrapper Wrapper[K, V]) Get(key K) V {
	return wrapper.m.Get(key)
}

func (wrapper Wrapper[K, V]) TryGet(key K) (V, bool) {
	return wrapper.m.TryGet(key)
}

func (wrapper Wrapper[K, V]) Keys() []K {
	return wrapper.m.Keys()
}

func (wrapper Wrapper[K, V]) Contains(key K) bool {
	return wrapper.m.Contains(key)
}

func (wrapper Wrapper[K, V]) Clone() Wrapper[K, V] {
	clonedMap := wrapper.m.Clone()
	return Wrapper[K, V]{clonedMap}
}

func (wrapper Wrapper[K, V]) Iterator() iterable.Iterator[misc.Pair[K, V]] {
	return wrapper.m.Iterator()
}

func (wrapper Wrapper[K, V]) Stream() streams.Stream[misc.Pair[K, V]] {
	return wrapper.m.Stream()
}
