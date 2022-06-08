package linkedlistmap

import (
	"github.com/djordje200179/GoExtendedLibrary/datastructures/maps"
	"github.com/djordje200179/GoExtendedLibrary/datastructures/sequences"
	"github.com/djordje200179/GoExtendedLibrary/misc"
)

type iterator[K comparable, V any] struct {
	sequences.Iterator[misc.Pair[K, V]]
}

func (it iterator[K, V]) Get() maps.Entry[K, V] {
	return it
}

func (it iterator[K, V]) GetKey() K {
	return it.Iterator.Get().First
}

func (it iterator[K, V]) GetValue() V {
	return it.Iterator.Get().Second
}

func (it iterator[K, V]) SetValue(value V) {
	data := it.Iterator.Get()
	data.Second = value
	it.Iterator.Set(data)
}
