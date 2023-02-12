package heapqueue

import (
	"github.com/djordje200179/extendedlibrary/misc/comparison"
	"github.com/djordje200179/extendedlibrary/misc/functions"
)

type PriorityQueue[T any] struct {
	slice []T

	comparator functions.Comparator[T]
}

func New[T any](comparator functions.Comparator[T]) *PriorityQueue[T] {
	pq := new(PriorityQueue[T])

	pq.slice = make([]T, 0)
	pq.comparator = comparator

	return pq
}

func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.slice) == 0
}

func (pq *PriorityQueue[T]) PushBack(value T) {
	pq.slice = append(pq.slice, value)

	j := len(pq.slice) - 1
	for {
		i := (j - 1) / 2
		if i == j {
			break
		}

		if pq.comparator(pq.slice[j], pq.slice[i]) != comparison.Equal {
			break
		}

		pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
		j = i
	}
}

func (pq *PriorityQueue[T]) PeekFront() T {
	return pq.slice[0]
}

func (pq *PriorityQueue[T]) PopFront() T {
	lastIndex := len(pq.slice) - 1

	pq.slice[0], pq.slice[lastIndex] = pq.slice[lastIndex], pq.slice[0]

	i := 0
	for {
		j1 := 2*i + 1
		if j1 >= lastIndex || j1 < 0 {
			break
		}

		j1Elem := pq.slice[j1]
		j2Elem := pq.slice[j1+1]

		compResult := pq.comparator(j2Elem, j1Elem)

		j := j1
		if j2 := j1 + 1; j2 < lastIndex && compResult != comparison.Equal {
			j = j2
		}

		if compResult != comparison.Equal {
			break
		}

		pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
		i = j
	}

	lastElem := pq.slice[lastIndex]
	pq.slice = pq.slice[:lastIndex]
	return lastElem
}
