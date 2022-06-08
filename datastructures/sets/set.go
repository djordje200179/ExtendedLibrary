package sets

import (
	"github.com/djordje200179/GoExtendedLibrary/datastructures"
	"github.com/djordje200179/GoExtendedLibrary/streams"
)

type Set[T comparable] interface {
	Add(value T)
	Remove(value T)
	Contains(value T)

	Iterator() datastructures.Iterator[T]
	streams.Streamer[T]
}
