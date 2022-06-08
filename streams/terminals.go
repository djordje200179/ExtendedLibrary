package streams

import "github.com/djordje200179/GoExtendedLibrary/misc/functions"

func (stream Stream[T]) ForEach(function func(curr T)) {
	for {
		data, ok := stream.getNext()
		if !ok {
			break
		}

		function(data)
	}
}

func (stream Stream[T]) Any(tester func(curr T) bool) bool {
	for {
		data, ok := stream.getNext()
		if !ok {
			break
		}

		if tester(data) {
			stream.signal <- end
			return true
		}
	}

	return false
}

func (stream Stream[T]) All(tester func(curr T) bool) bool {
	for {
		data, ok := stream.getNext()
		if !ok {
			break
		}

		if !tester(data) {
			stream.signal <- end
			return false
		}
	}

	return true
}

func (stream Stream[T]) Count() int {
	count := 0

	for {
		_, ok := stream.getNext()
		if !ok {
			break
		}

		count++
	}

	return count
}

func (stream Stream[T]) Max(less functions.Less[T]) T {
	var max T
	set := false

	for {
		data, ok := stream.getNext()
		if !ok {
			break
		}

		if !set || less(max, data) {
			max = data
			set = true
		}
	}

	return max
}

func (stream Stream[T]) Min(less functions.Less[T]) T {
	var min T
	set := false

	for {
		data, ok := stream.getNext()
		if !ok {
			break
		}

		if !set || less(data, min) {
			min = data
			set = true
		}
	}

	return min
}
