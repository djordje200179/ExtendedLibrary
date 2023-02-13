package bst

type iterator[K comparable, V any] struct {
	tree *BinarySearchTree[K, V]

	curr *node[K, V]
}

func (it *iterator[K, V]) Valid() bool {
	return it.curr != nil
}

func (it *iterator[K, V]) Move() {
	it.curr = it.curr.next()
}

func (it *iterator[K, V]) Get() K {
	return it.curr.key
}

func (it *iterator[K, V]) Value() V {
	return *it.ValueRef()
}

func (it *iterator[K, V]) ValueRef() *V {
	return &it.curr.value
}

func (it *iterator[K, V]) SetValue(value V) {
	*it.ValueRef() = value
}

func (it *iterator[K, V]) Remove() {
	it.tree.removeNode(it.curr)
}
