package linkedmap

import "github.com/djordje200179/extendedlibrary/misc"

type Iterator[K, V any] struct {
	wrapper *Wrapper[K, V]

	curr *Node[K, V]
}

func (it *Iterator[K, V]) Valid() bool {
	return it.curr != nil
}

func (it *Iterator[K, V]) Move() {
	if it.curr == nil {
		return
	}

	it.curr = it.curr.next
}

func (it *Iterator[K, V]) Get() misc.Pair[K, V] {
	return misc.Pair[K, V]{
		First:  it.Key(),
		Second: it.Value(),
	}
}

func (it *Iterator[K, V]) Key() K {
	return it.curr.key
}

func (it *Iterator[K, V]) Value() V {
	return it.curr.Value
}

func (it *Iterator[K, V]) ValueRef() *V {
	return &it.curr.Value
}

func (it *Iterator[K, V]) SetValue(value V) {
	it.curr.Value = value
}

func (it *Iterator[K, V]) Remove() {
	next := it.curr.next
	it.wrapper.m.Remove(it.curr.key)
	it.curr = next
}

func (it *Iterator[K, V]) Node() *Node[K, V] {
	return it.curr
}
