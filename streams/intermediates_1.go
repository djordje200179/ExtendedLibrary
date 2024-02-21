package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"slices"
)

func (s Stream[T]) Enumerate() Stream2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for elem := range s {
			if !yield(i, elem) {
				break
			}

			i++
		}
	}
}

func Map[T, U any](s Stream[T], mapper functions.Mapper[T, U]) Stream[U] {
	return func(yield func(U) bool) {
		for elem := range s {
			if !yield(mapper(elem)) {
				break
			}
		}
	}
}

func (s Stream[T]) Map(mapper functions.Mapper[T, T]) Stream[T] {
	return Map(s, mapper)
}

func (s Stream[T]) Filter(predicate predication.Predicate[T]) Stream[T] {
	return func(yield func(T) bool) {
		for elem := range s {
			if !predicate(elem) {
				continue
			}

			if !yield(elem) {
				break
			}
		}
	}
}

func (s Stream[T]) Limit(count int) Stream[T] {
	return func(yield func(T) bool) {
		i := 0
		for elem := range s {
			if i >= count {
				break
			}

			if !yield(elem) {
				break
			}

			i++
		}
	}
}

func (s Stream[T]) Seek(count int) Stream[T] {
	return func(yield func(T) bool) {
		i := 0
		for elem := range s {
			if i < count {
				i++
				continue
			}

			if !yield(elem) {
				break
			}

			i++
		}
	}
}

func (s Stream[T]) Sort(comparator comparison.Comparator[T]) Stream[T] {
	return func(yield func(T) bool) {
		var slice []T
		for elem := range s {
			slice = append(slice, elem)
		}

		slices.SortFunc(slice, comparator)

		for _, elem := range slice {
			if !yield(elem) {
				break
			}
		}
	}
}

func Window[T any](s Stream[T], width int) Stream[[]T] {
	return func(yield func([]T) bool) {
		window := make([]T, width)

		i := 0
		for elem := range s {
			if i < width {
				window[i] = elem
				i++
				continue
			}

			if !yield(window) {
				break
			}

			for i := range width - 1 {
				window[i] = window[i+1]
			}

			window[width-1] = elem
		}
	}
}
