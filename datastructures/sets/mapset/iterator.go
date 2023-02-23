package mapset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type iterator[T comparable] struct {
	maps.Iterator[T, empty]
}

func (it iterator[T]) Get() T {
	return it.Iterator.Key()
}
