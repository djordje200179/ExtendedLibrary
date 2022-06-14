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
				elem := stream.getNext()
				if !elem.IsPresent() {
					goto end
				}

				data := elem.GetOrPanic()
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

			elem := stream.getNext()
			if !elem.IsPresent() {
				break
			}

			ret.data <- elem.GetOrPanic()
		}

		ret.close()
	}()

	return ret
}

func (stream Stream[T]) Seek(count int) Stream[T] {
	ret := create[T]()

	go func() {
		seeked := false
		for ret.waitRequest() {
			if !seeked {
				for i := 0; i < count; i++ {
					elem := stream.getNext()
					if !elem.IsPresent() {
						goto end
					}
				}

				seeked = true
			}

			elem := stream.getNext()
			if !elem.IsPresent() {
				break
			}

			ret.data <- elem.GetOrPanic()
		}

	end:
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
