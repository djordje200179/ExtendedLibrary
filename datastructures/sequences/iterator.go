package sequences

import "github.com/djordje200179/extendedlibrary/datastructures"

type Iterator[T any] interface {
	datastructures.Iterator[T]

	Set(value T)

	InsertBefore(value T)
	InsertAfter(value T)

	Remove()
}
