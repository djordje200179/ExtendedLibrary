package bst

type entry[K comparable, V any] struct {
	node *node[K, V]
}

func (e entry[K, V]) Key() K {
	return e.node.key
}

func (e entry[K, V]) ValueRef() *V {
	return &e.node.value
}

func (e entry[K, V]) Value() V {
	return *e.ValueRef()
}

func (e entry[K, V]) SetValue(value V) {
	*e.ValueRef() = value
}