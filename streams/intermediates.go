package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams/suppliers"
	"sort"
)

func Map[T, U any](stream Stream[T], mapper functions.Mapper[T, U]) Stream[U] {
	generator := func() optional.Optional[U] {
		if elem := stream.supplier.Supply(); elem.HasValue() {
			return optional.FromValue(mapper(elem.Get()))
		} else {
			return optional.Empty[U]()
		}
	}

	supplier := suppliers.FromFiniteGenerator(generator)
	return New(supplier)
}

func (stream Stream[T]) Filter(predictor functions.Predictor[T]) Stream[T] {
	generator := func() optional.Optional[T] {
		for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
			data := elem.Get()

			if predictor(data) {
				return optional.FromValue(data)
			}
		}

		return optional.Empty[T]()
	}

	supplier := suppliers.FromFiniteGenerator(generator)
	return New(supplier)
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

	supplier := suppliers.FromFiniteGenerator(generator)
	return New(supplier)
}

func (stream Stream[T]) Seek(count int) Stream[T] {
	generator := func() optional.Optional[T] {
		for ; count > 0; count-- {
			stream.supplier.Supply()
		}

		return stream.supplier.Supply()
	}

	supplier := suppliers.FromFiniteGenerator(generator)
	return New(supplier)
}

func (stream Stream[T]) Sort(comparator functions.Comparator[T]) Stream[T] {
	var sortedSlice []T
	var sortedSupplier suppliers.Supplier[T]

	generator := func() optional.Optional[T] {
		if sortedSlice == nil {
			for elem := stream.supplier.Supply(); elem.HasValue(); elem = stream.supplier.Supply() {
				sortedSlice = append(sortedSlice, elem.Get())
			}

			sort.SliceStable(5, func(i, j int) bool { return comparator(sortedSlice[i], sortedSlice[j]) == comparison.FirstSmaller })

			sortedSupplier = suppliers.FromSlice(sortedSlice)
		}

		return sortedSupplier.Supply()
	}

	supplier := suppliers.FromFiniteGenerator(generator)
	return New(supplier)
}
