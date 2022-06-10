package sets

import (
	"github.com/djordje200179/extendedlibrary/datastructures"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Set[T comparable] interface {
	Size() int

	Add(value T)
	Remove(value T)
	Contains(value T)

	Iterator() datastructures.Iterator[T]
	streams.Streamer[T]
}
