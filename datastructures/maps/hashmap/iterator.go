package hashmap

type iterator[K comparable, V any] struct {
	m     HashMap[K, V]
	keys  []K
	index int
}

func (it *iterator[K, V]) Valid() bool {
	return it.index < len(it.keys)
}

func (it *iterator[K, V]) Move() {
	it.index++
}

func (it *iterator[K, V]) Get() K {
	return it.keys[it.index]
}

func (it *iterator[K, V]) Value() V {
	return it.m.Get(it.keys[it.index])
}

func (it *iterator[K, V]) ValueRef() *V {
	return it.m.GetRef(it.keys[it.index])
}

func (it *iterator[K, V]) SetValue(value V) {
	it.m.Set(it.keys[it.index], value)
}

func (it *iterator[K, V]) Remove() {
	it.m.Remove(it.keys[it.index])
}
