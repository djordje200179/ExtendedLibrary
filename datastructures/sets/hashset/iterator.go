package hashset

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type iterator[T comparable] struct {
	datastructures.Iterator[maps.Entry[T, bool]]
}

func (it iterator[T]) Get() T { return it.Iterator.Get().Key() }
