package readonlyset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/sets"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Wrapper[T any] struct {
	set sets.Set[T]
}

func From[T any](set sets.Set[T]) Wrapper[T] {
	return Wrapper[T]{set}
}

func (wrapper Wrapper[T]) Size() int {
	return wrapper.set.Size()
}

func (wrapper Wrapper[T]) Contains(value T) bool {
	return wrapper.set.Contains(value)
}

func (wrapper Wrapper[T]) Clone() Wrapper[T] {
	clonedSet := wrapper.set.Clone()
	return Wrapper[T]{clonedSet}
}

func (wrapper Wrapper[T]) Iterator() iterable.Iterator[T] {
	return wrapper.set.Iterator()
}

func (wrapper Wrapper[T]) Stream() streams.Stream[T] {
	return wrapper.set.Stream()
}
