package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"sort"
)

func Map[T, U any](stream Stream[T], mapper functions.Mapper[T, U]) Stream[U] {
	generator := func() optional.Optional[U] {
		if elem := stream.supplier(); elem.Valid {
			return optional.FromValue(mapper(elem.Value))
		} else {
			return optional.Empty[U]()
		}
	}

	return Stream[U]{generator}
}

func (stream Stream[T]) Map(mapper functions.Mapper[T, T]) Stream[T] {
	return Map(stream, mapper)
}

func (stream Stream[T]) Filter(predicate predication.Predicate[T]) Stream[T] {
	generator := func() optional.Optional[T] {
		for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
			if predicate(elem.Value) {
				return optional.FromValue(elem.Value)
			}
		}

		return optional.Empty[T]()
	}

	return Stream[T]{generator}
}

func (stream Stream[T]) Limit(count int) Stream[T] {
	generator := func() optional.Optional[T] {
		if count > 0 {
			count--
			return stream.supplier()
		} else {
			return optional.Empty[T]()
		}
	}

	return Stream[T]{generator}
}

func (stream Stream[T]) Seek(count int) Stream[T] {
	generator := func() optional.Optional[T] {
		for ; count > 0; count-- {
			stream.supplier()
		}

		return stream.supplier()
	}

	return Stream[T]{generator}
}

func (stream Stream[T]) Sort(comparator comparison.Comparator[T]) Stream[T] {
	var index int
	var sortedSlice []T

	generator := func() optional.Optional[T] {
		if sortedSlice == nil {
			for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
				sortedSlice = append(sortedSlice, elem.Value)
			}

			sort.SliceStable(sortedSlice, func(i, j int) bool {
				return comparator(sortedSlice[i], sortedSlice[j]) == comparison.FirstSmaller
			})
		}

		if index < len(sortedSlice) {
			value := sortedSlice[index]
			index++
			return optional.FromValue(value)
		} else {
			return optional.Empty[T]()
		}
	}

	return Stream[T]{generator}
}

func Window[T any](stream Stream[T], width int) Stream[[]T] {
	var window []T

	generator := func() optional.Optional[[]T] {
		if window == nil {
			window = make([]T, width)
			i := 0
			for ; i < width; i++ {
				elem := stream.supplier()
				if !elem.Valid {
					break
				}

				window[i] = elem.Value
			}

			if i == width {
				return optional.FromValue(window)
			} else {
				return optional.Empty[[]T]()
			}
		} else {
			for i := range width - 1 {
				window[i] = window[i+1]
			}

			if elem := stream.supplier(); elem.Valid {
				window[width-1] = elem.Value
				return optional.FromValue(window)
			} else {
				return optional.Empty[[]T]()
			}

		}
	}

	return Stream[[]T]{generator}
}
