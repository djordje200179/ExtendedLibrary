package hashmap

type entry[K comparable, V any] struct {
	m   HashMap[K, V]
	key K
}

func (e entry[K, V]) Key() K { return e.key }

func (e entry[K, V]) ValueRef() *V     { panic("Not supported") }
func (e entry[K, V]) Value() V         { return e.m[e.key] }
func (e entry[K, V]) SetValue(value V) { e.m[e.key] = value }
