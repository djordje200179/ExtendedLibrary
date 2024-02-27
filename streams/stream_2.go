package streams

import "iter"

type Stream2[K, V any] iter.Seq2[K, V]

type Streamer2[K, V any] interface {
	Stream2(yield func(K, V) bool)
}

func From2[K, V any](streamer Streamer2[K, V]) Stream2[K, V] {
	return streamer.Stream2
}

func FromMap[K comparable, V any](m map[K]V) Stream2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}

func Zip[K, V any](s1 Stream[K], s2 Stream[V]) Stream2[K, V] {
	return func(yield func(K, V) bool) {
		for k := range s1 {
			for v := range s2 {
				if !yield(k, v) {
					return
				}

				break
			}
		}
	}
}
