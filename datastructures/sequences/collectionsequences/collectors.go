package collectionsequences

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/streams"
)

func QueueCollector[T any]() streams.Collector[T, sequences.Queue[T]] {
	return sequences.Collector[T, sequences.Queue[T]]{
		BackPusher: NewDeque[T](),
	}
}

func StackCollector[T any]() streams.Collector[T, sequences.Stack[T]] {
	return sequences.Collector[T, sequences.Stack[T]]{
		BackPusher: NewDeque[T](),
	}
}
