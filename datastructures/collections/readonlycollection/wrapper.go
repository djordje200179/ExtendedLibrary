package readonlycollection

import (
	"github.com/djordje200179/extendedlibrary/datastructures/collections"
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Wrapper[T any] struct {
	collection collections.Collection[T]
}

func From[T any](collection collections.Collection[T]) Wrapper[T] {
	return Wrapper[T]{collection}
}

func (wrapper Wrapper[T]) Size() int {
	return wrapper.collection.Size()
}

func (wrapper Wrapper[T]) Get(index int) T {
	return wrapper.collection.Get(index)
}

func (wrapper Wrapper[T]) Clone() Wrapper[T] {
	clonedCollection := wrapper.collection.Clone()
	return Wrapper[T]{clonedCollection}
}

func (wrapper Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.collection.Iterator()
}

func (wrapper Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.collection.Stream()
}

func (wrapper *Wrapper[T]) FindIndex(predicate predication.Predicate[T]) (int, bool) {
	return wrapper.collection.FindIndex(predicate)
}
