package streams

import "github.com/djordje200179/extendedlibrary/misc/functions/predication"

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

func Window2[K, V any](s Stream2[K, V], size int) Stream2[K, []V] {
	return func(yield func(K, []V) bool) {
		window := make([]V, size)
		i := 0
		for k, v := range s {
			if i < size {
				window[i] = v
				i++
				continue
			}

			if !yield(k, window) {
				break
			}

			copy(window, window[1:])
			window[size-1] = v
		}
	}
}
