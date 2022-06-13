package maps

import "github.com/djordje200179/extendedlibrary/datastructures"

type Iterator[K comparable, V any] interface {
	datastructures.Iterator[Entry[K, V]]

	Remove()
}
