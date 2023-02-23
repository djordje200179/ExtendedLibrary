package linkedlist

type Iterator[T any] struct {
	list *List[T]

	curr  *Node[T]
	index int
}

func (it *Iterator[T]) Valid() bool {
	return it.curr != nil
}

func (it *Iterator[T]) Move() {
	it.curr = it.curr.next
	it.index++
}

func (it *Iterator[T]) GetRef() *T {
	return &it.curr.Value
}

func (it *Iterator[T]) Get() T {
	return it.curr.Value
}

func (it *Iterator[T]) Set(value T) {
	it.curr.Value = value
}

func (it *Iterator[T]) InsertBefore(value T) {
	it.list.insertBefore(it.curr, value)
	it.index++
}

func (it *Iterator[T]) InsertAfter(value T) {
	it.list.insertAfter(it.curr, value)
}

func (it *Iterator[T]) Remove() {
	it.list.removeNode(it.curr)
}

func (it *Iterator[T]) Index() int {
	return it.index
}
