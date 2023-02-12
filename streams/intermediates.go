package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"sort"
)

func Map[T, U any](stream Stream[T], mapper functions.Mapper[T, U]) Stream[U] {
	generator := func() optional.Optional[U] {
		if elem := stream.supplier.Supply(); elem.Valid {
			return optional.FromValue(mapper(elem.Value))
		} else {
			return optional.Empty[U]()
		}
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Filter(predictor functions.Predictor[T]) Stream[T] {
	generator := func() optional.Optional[T] {
		for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
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
			return stream.supplier.Supply()
		} else {
			return optional.Empty[T]()
		}
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Seek(count int) Stream[T] {
	generator := func() optional.Optional[T] {
		for ; count > 0; count-- {
			stream.supplier.Supply()
		}

		return stream.supplier.Supply()
	}

	return FromFiniteGenerator(generator)
}

func (stream Stream[T]) Sort(comparator functions.Comparator[T]) Stream[T] {
	var sortedSlice []T
	var sortedSupplier suppliers.Supplier[T]

	generator := func() optional.Optional[T] {
		if sortedSlice == nil {
			for elem := stream.supplier.Supply(); elem.Valid; elem = stream.supplier.Supply() {
				sortedSlice = append(sortedSlice, elem.Value)
			}

			sort.SliceStable(5, func(i, j int) bool { return comparator(sortedSlice[i], sortedSlice[j]) == comparison.FirstSmaller })

			sortedSupplier = suppliers.Slice(sortedSlice)
		}

		return sortedSupplier.Supply()
	}

	return FromFiniteGenerator(generator)
}
