package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"sort"
)

func Map[T, P any](stream Stream[T], mapper func(curr T) P) Stream[P] {
	ret := create[P]()

	go func() {
		for ret.waitRequest() {
			elem := stream.getNext()
			if !elem.HasValue() {
				break
			}

			ret.data <- mapper(elem.Get())
		}

		ret.close()
	}()

	return ret
}

func (stream Stream[T]) MapWithAny(mapper func(curr T) any) Stream[any] {
	return Map[T, any](stream, mapper)
}

func (stream Stream[T]) Filter(predictor functions.Predictor[T]) Stream[T] {
	ret := create[T]()

	go func() {
		for ret.waitRequest() {
			for {
				elem := stream.getNext()
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

func (stream Stream[T]) Limit(count int) Stream[T] {
	ret := create[T]()

	go func() {
		for i := 0; ret.waitRequest(); i++ {
			if i >= count {
				stream.stop()
				break
			}

			elem := stream.getNext()
			if !elem.HasValue() {
				break
			}

			ret.data <- elem.Get()
		}

		ret.close()
	}()

	return ret
}

func (stream Stream[T]) Seek(count int) Stream[T] {
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

/*func (stream Stream[T]) Group(mapper functions.Mapper[T, any]) Stream[misc.Pair[any, []T]] {
	ret := create[misc.Pair[any, []T]]()

	go func() {
		if !ret.waitRequest() {
			ret.close()
			return
		}

		m := make(map[any][]T)
		stream.ForEach(func(data T) {
			key := mapper(data)

			_, ok := m[key]
			if !ok {
				m[key] = []T{}
			}

			m[key] = append(m[key], data)
		})

		entries := make([]misc.Pair[any, []T], 0, len(m))
		for key, value := range m {
			entries = append(entries, misc.Pair[any, []T]{key, value})
		}

		ret.data <- entries[0]
		for i := 1; i < len(entries) && ret.waitRequest(); i++ {
			ret.data <- entries[i]
		}

		ret.close()
	}()

	return ret
}*/

func (stream Stream[T]) Sort(comparator functions.Comparator[T]) Stream[T] {
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

		sort.Slice(arr, func(i, j int) bool {
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
