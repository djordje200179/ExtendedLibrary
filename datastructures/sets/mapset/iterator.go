package mapset

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps"
)

type Iterator[T any] struct {
	mapIt maps.Iterator[T, empty]
}

func (it Iterator[T]) Valid() bool {
	return it.mapIt.Valid()
}

func (it Iterator[T]) Move() {
	it.mapIt.Move()
}

func (it Iterator[T]) Get() T {
	return it.mapIt.Key()
}

func (it Iterator[T]) Remove() {
	it.mapIt.Remove()
}
