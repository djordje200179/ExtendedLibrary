package maps

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
)

type Entry[K comparable, V any] interface {
	GetKey() K

	GetValue() V
	SetValue(value V)

	Remove()
}

type Iterator[K comparable, V any] interface {
	datastructures.Iterator[Entry[K, V]]
}
