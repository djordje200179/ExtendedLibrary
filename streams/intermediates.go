package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"sort"
)

func Map[T, P any](stream *Stream[T], mapper func(curr T) P) *Stream[P] {
	ret := create[P]()

	go func() {
		for elem := stream.getNext(); elem.HasValue() && ret.waitRequest(); elem = stream.getNext() {
			ret.data <- mapper(elem.Get())
		}

		ret.close()
	}()

	return ret
}

func (stream *Stream[T]) MapWithAny(mapper func(curr T) any) *Stream[any] {
	return Map[T, any](stream, mapper)
}

func (stream *Stream[T]) Filter(predictor functions.Predictor[T]) *Stream[T] {
	ret := create[T]()

	go func() {
		for ret.waitRequest() {
			for elem := stream.getNext(); ; elem = stream.getNext() {
				if !elem.HasValue() {
					goto end
				}

				data := elem.Get()
				if predictor(data) {
					ret.data <- data
					break
				}
			}
		}

	end:
		ret.close()
	}()

	return ret
}

func (stream *Stream[T]) Limit(count int) *Stream[T] {
	ret := create[T]()

	go func() {
		i := 0
		for elem := stream.getNext(); i < count && elem.HasValue() && ret.waitRequest(); i, elem = i+1, stream.getNext() {
			ret.data <- elem.Get()
		}

		if i == count {
			stream.stop()
		}

		ret.close()
	}()

	return ret
}

func (stream *Stream[T]) Seek(count int) *Stream[T] {
	ret := create[T]()

	go func() {
		if !ret.waitRequest() {
			stream.stop()
			return
		}

		for i := 0; i <= count; i++ {
			elem := stream.getNext()
			if !elem.HasValue() {
				goto end
			}

			if i == count {
				ret.data <- elem.Get()
			}
		}

		for ret.waitRequest() {
			elem := stream.getNext()
			if !elem.HasValue() {
				break
			}

			ret.data <- elem.Get()
		}

	end:
		ret.close()
	}()

	return ret
}

func (stream *Stream[T]) Sort(comparator functions.Comparator[T]) *Stream[T] {
	ret := create[T]()

	go func() {
		if !ret.waitRequest() {
			ret.close()
			return
		}

		arr := make([]T, 0)
		stream.ForEach(func(data T) {
			arr = append(arr, data)
		})

		sort.SliceStable(arr, func(i, j int) bool {
			return comparator(arr[i], arr[j]) == comparison.FirstSmaller
		})

		ret.data <- arr[0]
		for i := 1; i < len(arr) && ret.waitRequest(); i++ {
			ret.data <- arr[i]
		}

		ret.close()
	}()

	return ret
}
