package streams

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/array"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"sort"
)

type Mapper[T, P any] func(value T) P

func (stream Stream[T]) Map(mapper func(curr T) any) Stream[any] {
	ret := create[any]()

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

func (stream Stream[T]) Filter(predicate functions.Predicate[T]) Stream[T] {
	ret := create[T]()

	go func() {
		for ret.waitRequest() {
			for {
				elem := stream.getNext()
				if !elem.HasValue() {
					goto end
				}

				data := elem.Get()
				if predicate(data) {
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
		seeked := false
		for ret.waitRequest() {
			if !seeked {
				for i := 0; i < count; i++ {
					elem := stream.getNext()
					if !elem.HasValue() {
						goto end
					}
				}

				seeked = true
			}

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

func (stream Stream[T]) Group(mapper Mapper[T, any]) Stream[misc.Pair[any, sequences.Sequence[T]]] {
	ret := create[misc.Pair[any, sequences.Sequence[T]]]()

	go func() {
		if !ret.waitRequest() {
			ret.close()
			return
		}

		m := hashmap.New[any, sequences.Sequence[T]]()
		stream.ForEach(func(data T) {
			key := mapper(data)

			if !m.Contains(key) {
				m.Set(key, array.New[T](0))
			}

			m.Get(key).Append(data)
		})

		it := m.Iterator()
		ret.data <- it.Get().Get()
		it.Move()

		for ; it.IsValid() && ret.waitRequest(); it.Move() {
			ret.data <- it.Get().Get()
		}

		ret.close()
	}()

	return ret
}

func (stream Stream[T]) Sort(comparator comparison.Comparator[T]) Stream[T] {
	ret := create[T]()

	go func() {
		var arr []T

		for i := 0; (arr == nil || i < len(arr)) && ret.waitRequest(); i++ {
			if arr == nil {
				arr = make([]T, 0)

				stream.ForEach(func(data T) {
					arr = append(arr, data)
				})

				sort.Slice(arr, func(i, j int) bool {
					return comparator(arr[i], arr[j]) == comparison.FirstSmaller
				})
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
