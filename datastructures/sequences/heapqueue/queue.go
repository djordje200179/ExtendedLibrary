package heapqueue

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Queue[T any] struct {
	slice []T

	comparator functions.Comparator[T]
}

func New[T any](comparator functions.Comparator[T]) *Queue[T] {
	pq := new(Queue[T])

	pq.slice = make([]T, 0)
	pq.comparator = comparator

	return pq
}

func Collector[T any](comparator functions.Comparator[T]) streams.Collector[T, *Queue[T]] {
	return sequences.Collector[T, *Queue[T]]{
		BackPusher: New[T](comparator),
	}
}

func (pq *Queue[T]) Empty() bool {
	return len(pq.slice) == 0
}

func (pq *Queue[T]) PushBack(value T) {
	pq.slice = append(pq.slice, value)

	j := len(pq.slice) - 1
	for {
		i := (j - 1) / 2
		if i == j || pq.comparator(pq.slice[j], pq.slice[i]) != comparison.FirstSmaller {
			break
		}

		pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
		j = i
	}
}

func (pq *Queue[T]) PeekFront() T {
	return pq.slice[0]
}

func (pq *Queue[T]) PopFront() T {
	lastIndex := len(pq.slice) - 1

	pq.slice[0], pq.slice[lastIndex] = pq.slice[lastIndex], pq.slice[0]

	i := 0
	for {
		j1 := 2*i + 1
		if j1 >= lastIndex || j1 < 0 {
			break
		}

		j := j1
		if j2 := j1 + 1; j2 < lastIndex && pq.comparator(pq.slice[j2], pq.slice[j1]) == comparison.FirstSmaller {
			j = j2
		}

		if pq.comparator(pq.slice[j], pq.slice[i]) != comparison.FirstSmaller {
			break
		}

		pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
		i = j
	}

	lastElem := pq.slice[lastIndex]
	pq.slice = pq.slice[:lastIndex]
	return lastElem
}
