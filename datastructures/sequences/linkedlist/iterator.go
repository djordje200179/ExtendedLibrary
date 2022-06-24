package linkedlist

type Iterator[T any] struct {
	list *LinkedList[T]
	curr *Node[T]
}

func (it *Iterator[T]) Valid() bool { return it.curr != nil }
func (it *Iterator[T]) Move()       { it.curr = it.curr.next }

func (it *Iterator[T]) GetRef() *T  { return &it.curr.Value }
func (it *Iterator[T]) Get() T      { return it.curr.Value }
func (it *Iterator[T]) Set(value T) { it.curr.Value = value }

func (it *Iterator[T]) InsertBefore(value T) { it.list.insertBeforeNode(it.curr, value) }
func (it *Iterator[T]) InsertAfter(value T)  { it.list.insertAfterNode(it.curr, value) }
func (it *Iterator[T]) Remove()              { it.list.removeNode(it.curr) }

func (it *Iterator[T]) Node() *Node[T] { return it.curr }
