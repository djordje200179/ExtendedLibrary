package array

type iterator[T any] struct {
	array *Array[T]
	index int
}

func (it *iterator[T]) IsValid() bool {
	return it.index < it.array.Size()
}

func (it *iterator[T]) Move() {
	it.index++
}

func (it *iterator[T]) Get() T {
	return it.array.Get(it.index)
}

func (it *iterator[T]) Set(value T) {
	it.array.Set(it.index, value)
}

func (it *iterator[T]) InsertBefore(value T) {
	it.array.Insert(it.index, value)
}

func (it *iterator[T]) InsertAfter(value T) {
	it.array.Insert(it.index+1, value)
}

func (it *iterator[T]) Remove() {
	it.array.Remove(it.index)
}
