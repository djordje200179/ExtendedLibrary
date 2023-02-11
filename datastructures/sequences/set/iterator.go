package set

import (
	"github.com/djordje200179/extendedlibrary/datastructures/iterable"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type iterator[T comparable] struct {
	iterable.Iterator[maps.Entry[T, empty]]
}

func (it iterator[T]) Get() T { return it.Iterator.Get().Key() }
