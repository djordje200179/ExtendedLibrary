package synclist

type Iterator[T any] struct {
	curr *Node[T]
}

func (it *Iterator[T]) Valid() bool {
	return it.curr != nil
}

func (it *Iterator[T]) Move() {
	it.curr = it.curr.next
}

func (it *Iterator[T]) Get() T {
	return it.curr.Value
}
