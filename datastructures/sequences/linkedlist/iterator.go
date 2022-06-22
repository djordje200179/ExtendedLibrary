package linkedlist

type Iterator[T any] struct {
	list    *LinkedList[T]
	current *Node[T]
}

func (it *Iterator[T]) Valid() bool { return it.current != nil }
func (it *Iterator[T]) Move()       { it.current = it.current.next }

func (it *Iterator[T]) Get() T      { return it.current.Value }
func (it *Iterator[T]) Set(value T) { it.current.Value = value }

func (it *Iterator[T]) InsertBefore(value T) { it.list.insertBeforeNode(it.current, value) }
func (it *Iterator[T]) InsertAfter(value T)  { it.list.insertAfterNode(it.current, value) }
func (it *Iterator[T]) Remove()              { it.list.removeNode(it.current) }
