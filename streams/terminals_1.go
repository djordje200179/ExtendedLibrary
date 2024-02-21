package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

func (s Stream[T]) ForEach(f functions.ParamCallback[T]) {
	for elem := range s {
		f(elem)
	}
}

type Reducer[T, P any] func(acc P, value T) P

func Reduce[T, P any](s Stream[T], acc P, reducer Reducer[T, P]) P {
	for elem := range s {
		acc = reducer(acc, elem)
	}

	return acc
}

func (s Stream[T]) Reduce(acc T, reducer Reducer[T, T]) T {
	return Reduce(s, acc, reducer)
}

func (s Stream[T]) Any(predicate predication.Predicate[T]) bool {
	for elem := range s {
		if predicate(elem) {
			return true
		}
	}

	return false
}

func (s Stream[T]) All(predicate predication.Predicate[T]) bool {
	for elem := range s {
		if !predicate(elem) {
			return false
		}
	}

	return true
}

type Collector[T, R any] interface {
	Supply(value T)
	Finish() R
}

func Collect[T, R any](s Stream[T], collector Collector[T, R]) R {
	s.ForEach(collector.Supply)

	return collector.Finish()
}

func (s Stream[T]) Collect(c Collector[T, T]) T {
	return Collect(s, c)
}

func (s Stream[T]) Count() int {
	count := 0
	for _ = range s {
		count++
	}

	return count
}

func (s Stream[T]) Max(comparator comparison.Comparator[T]) (T, bool) {
	var currMax T
	set := false

	for elem := range s {
		if !set || comparator(elem, currMax) == comparison.FirstBigger {
			currMax = elem
			set = true
		}
	}

	return currMax, set
}

func (s Stream[T]) Min(comparator comparison.Comparator[T]) (T, bool) {
	var currMin T
	set := false

	for elem := range s {
		if !set || comparator(elem, currMin) == comparison.FirstSmaller {
			currMin = elem
			set = true
		}
	}

	return currMin, set
}

func (s Stream[T]) First() (T, bool) {
	for elem := range s {
		return elem, true
	}

	var zero T
	return zero, false
}

func (s Stream[T]) Find(predicate predication.Predicate[T]) (T, bool) {
	return s.Filter(predicate).First()
}

func (s Stream[T]) Channel() <-chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)

		for elem := range s {
			ch <- elem
		}
	}()

	return ch
}
