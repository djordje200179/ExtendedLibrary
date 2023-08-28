package array

type Iterator[T any] struct {
	array *Array[T]
	index int
}

func (it *Iterator[T]) Valid() bool {
	return it.index < it.array.Size()
}

func (it *Iterator[T]) Move() {
	it.index++
}

func (it *Iterator[T]) GetRef() *T {
	return it.array.GetRef(it.index)
}

func (it *Iterator[T]) Get() T {
	return it.array.Get(it.index)
}

func (it *Iterator[T]) Set(value T) {
	it.array.Set(it.index, value)
}

func (it *Iterator[T]) InsertBefore(value T) {
	it.array.Insert(it.index, value)
}

func (it *Iterator[T]) InsertAfter(value T) {
	it.array.Insert(it.index+1, value)
}

func (it *Iterator[T]) Remove() {
	it.array.Remove(it.index)
}

func (it *Iterator[T]) Index() int {
	return it.index
}
