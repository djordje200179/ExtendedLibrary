package hashmap

type entry[K comparable, V any] struct {
	m   Map[K, V]
	key K
}

func (e entry[K, V]) GetKey() K {
	return e.key
}

func (e entry[K, V]) GetValue() V {
	return e.m.Get(e.key)
}

func (e entry[K, V]) SetValue(value V) {
	e.m.Set(e.key, value)
}

func (e entry[K, V]) Remove() {
	e.m.Remove(e.key)
}
