package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"sort"
)

func Map[T, U any](stream Stream[T], mapper functions.Mapper[T, U]) Stream[U] {
	generator := func() optional.Optional[U] {
		if elem := stream.Supplier.Supply(); elem.Valid {
			return optional.FromValue(mapper(elem.Value))
		} else {
			return optional.Empty[U]()
		}
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Filter(predictor predication.Predictor[T]) Stream[T] {
	generator := func() optional.Optional[T] {
		for elem := stream.Supplier.Supply(); elem.Valid; elem = stream.Supplier.Supply() {
			if predictor(elem.Value) {
				return optional.FromValue(elem.Value)
			}
		}

		return optional.Empty[T]()
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Limit(count int) Stream[T] {
	generator := func() optional.Optional[T] {
		if count > 0 {
			count--
			return stream.Supplier.Supply()
		} else {
			return optional.Empty[T]()
		}
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Seek(count int) Stream[T] {
	generator := func() optional.Optional[T] {
		for ; count > 0; count-- {
			stream.Supplier.Supply()
		}

		return stream.Supplier.Supply()
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Sort(comparator comparison.Comparator[T]) Stream[T] {
	var sortedSlice []T
	var sortedSupplier suppliers.Supplier[T]

	generator := func() optional.Optional[T] {
		if sortedSlice == nil {
			for elem := stream.Supplier.Supply(); elem.Valid; elem = stream.Supplier.Supply() {
				sortedSlice = append(sortedSlice, elem.Value)
			}

			sort.SliceStable(sortedSlice, func(i, j int) bool {
				return comparator(sortedSlice[i], sortedSlice[j]) == comparison.FirstSmaller
			})

			sortedSupplier = suppliers.Slice(sortedSlice)
		}

		return sortedSupplier.Supply()
	}

	return FromFiniteGenerator(generator)
}
