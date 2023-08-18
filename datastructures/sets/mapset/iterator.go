package mapset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type Iterator[T any] struct {
	maps.Iterator[T, empty]
}

func (it Iterator[T]) Get() T {
	return it.Iterator.Key()
}
