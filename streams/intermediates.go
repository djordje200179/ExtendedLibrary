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

func (stream Stream[T]) Filter(predictor predication.Predictor[T]) Stream[T] {
	generator := func() optional.Optional[T] {
		for elem := stream.supplier(); elem.Valid; elem = stream.supplier() {
			if predictor(elem.Value) {
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
