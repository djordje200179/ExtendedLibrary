package collectionsequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/streams"
)

func Collector[T any]() streams.Collector[T, sequences.Queue[T]] {
	return sequences.Collector[T, sequences.Queue[T]]{
		BackPusher: NewDeque[T](),
	}
}
