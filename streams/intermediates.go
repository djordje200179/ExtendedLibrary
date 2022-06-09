package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"sort"
)

//func (stream Stream[T]) Map[P any](mapper func(curr T) P) Stream[P] {
//
//}

func (stream Stream[T]) Filter(predicate functions.Predicate[T]) Stream[T] {
	ret := create[T]()

	go func() {
		for ret.waitRequest() {
			for found := false; !found; {
				data, ok := stream.getNext().Get()
				if !ok {
					goto end
				}

				if predicate(data) {
					ret.data <- data
					found = true
				}
			}
		}

	end:
		ret.close()
	}()

	return ret
}

//func (stream Stream[T]) Reduce[P any](acc P, reducer func(acc P, curr T) P) Stream[P] {
//
//}

func (stream Stream[T]) Limit(count int) Stream[T] {
	ret := create[T]()

	go func() {
		for i := 0; ret.waitRequest(); i++ {
			if i >= count {
				stream.stop()
				break
			}

			data, ok := stream.getNext().Get()
			if !ok {
				break
			}

			ret.data <- data
		}

		ret.close()
	}()

	return ret
}

//func (stream Stream[T]) Group[P any](grouper func(curr T) P) Stream[misc.Pair[P, sequence.Sequence[T]] {
//
//}

func initArray[T any](arrPtr *[]T, stream Stream[T], comparator comparison.Comparator[T]) {
	*arrPtr = make([]T, 0)

	stream.ForEach(func(data T) {
		*arrPtr = append(*arrPtr, data)
	})

	arr := *arrPtr

	sort.Slice(*arrPtr, func(i, j int) bool {
		return comparator(arr[i], arr[j]) == comparison.FirstSmaller
	})
}

func (stream Stream[T]) Sort(comparator comparison.Comparator[T]) Stream[T] {
	ret := create[T]()

	go func() {
		var arr []T

		for i := 0; ret.waitRequest(); i++ {
			if arr == nil {
				initArray(&arr, stream, comparator)
			}

			if i >= len(arr) {
				break
			}

			ret.data <- arr[i]
		}

		ret.close()
	}()

	return ret
}
