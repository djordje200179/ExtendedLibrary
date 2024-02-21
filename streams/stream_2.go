package streams

import "iter"

type Stream2[K, V any] iter.Seq2[K, V]

type Streamer2[K, V any] interface {
	Stream2(yield func(K, V) bool)
}

func From2[K, V any](seq iter.Seq2[K, V]) Stream2[K, V] {
	return Stream2[K, V](seq)
}
