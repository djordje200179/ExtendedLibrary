package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

func (stream Stream[T]) ForEach(function func(curr T)) {
	for {
		data, ok := stream.getNext().Get()
		if !ok {
			break
		}

		function(data)
	}
}

func (stream Stream[T]) Any(predicate functions.Predicate[T]) bool {
	for {
		data, ok := stream.getNext().Get()
		if !ok {
			break
		}

		if predicate(data) {
			stream.stop()
			return true
		}
	}

	return false
}

func (stream Stream[T]) All(predicate functions.Predicate[T]) bool {
	for {
		data, ok := stream.getNext().Get()
		if !ok {
			break
		}

		if !predicate(data) {
			stream.stop()
			return false
		}
	}

	return true
}

func (stream Stream[T]) Count() int {
	count := 0

	for {
		_, ok := stream.getNext().Get()
		if !ok {
			break
		}

		count++
	}

	return count
}

func (stream Stream[T]) Max(comparator comparison.Comparator[T]) optional.Optional[T] {
	var max T
	set := false

	for {
		data, ok := stream.getNext().Get()
		if !ok {
			break
		}

		if !set || comparator(data, max) == comparison.FirstBigger {
			max = data
			set = true
		}
	}

	return optional.New(max, set)
}

func (stream Stream[T]) Min(comparator comparison.Comparator[T]) optional.Optional[T] {
	var min T
	set := false

	for {
		data, ok := stream.getNext().Get()
		if !ok {
			break
		}

		if !set || comparator(data, min) == comparison.FirstSmaller {
			min = data
			set = true
		}
	}

	return optional.New(min, set)
}

func (stream Stream[T]) First() optional.Optional[T] {
	data, ok := stream.getNext().Get()
	if !ok {
		return optional.Empty[T]()
	}

	stream.stop()
	return optional.FromValue(data)
}
