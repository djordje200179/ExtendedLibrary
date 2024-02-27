package streams

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

func (s Stream2[K, V]) Keys() Stream[K] {
	return func(yield func(K) bool) {
		for k, _ := range s {
			if !yield(k) {
				break
			}
		}
	}
}

func (s Stream2[K, V]) Values() Stream[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				break
			}
		}
	}
}

func (s Stream2[K, V]) Pairs() Stream[misc.Pair[K, V]] {
	return func(yield func(misc.Pair[K, V]) bool) {
		for k, v := range s {
			if !yield(misc.MakePair(k, v)) {
				break
			}
		}
	}
}

func (s Stream2[K, V]) FilterKeys(predicate predication.Predicate[K]) Stream2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range s {
			if !predicate(key) {
				continue
			}

			if !yield(key, value) {
				break
			}
		}
	}
}

func (s Stream2[K, V]) FilterValues(predicate predication.Predicate[V]) Stream2[K, V] {
	return func(yield func(K, V) bool) {
		for key, value := range s {
			if !predicate(value) {
				continue
			}

			if !yield(key, value) {
				break
			}
		}
	}
}

func (s Stream2[K, V]) Limit(count int) Stream2[K, V] {
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range s {
			if i >= count {
				break
			}

			if !yield(k, v) {
				break
			}

			i++
		}
	}
}

func (s Stream2[K, V]) Seek(count int) Stream2[K, V] {
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range s {
			if i < count {
				i++
				continue
			}

			if !yield(k, v) {
				break
			}
		}
	}
}
