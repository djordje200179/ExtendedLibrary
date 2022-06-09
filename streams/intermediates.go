package streams

//func (stream Stream[T]) Map[P any](mapper func(curr T) P) Stream[P] {
//
//}

func (stream Stream[T]) Filter(filter func(curr T) bool) Stream[T] {
	ret := create[T]()

	go func() {
		for <-ret.signal == next {
			for found := false; !found; {
				data, ok := stream.getNext().Get()
				if !ok {
					goto end
				}

				if filter(data) {
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
		for i := 0; <-ret.signal == next; i++ {
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

//func (stream Stream[T]) Sort(less functions.Less[T]) Stream[T] {
//
//}
