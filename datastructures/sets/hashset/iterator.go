package hashset

import (
	"github.com/djordje200179/GoExtendedLibrary/datastructures/maps"
)

type iterator[T comparable] struct {
	maps.Iterator[T, bool]
}

func (it iterator[T]) Get() T {
	return it.Iterator.Get().GetKey()
}
