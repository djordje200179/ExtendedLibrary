package priorityqueue

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
	pq := &Queue[T]{
		slice: make([]T, 0),

		comparator: comparator,
	}

	return pq
}

func Collector[T any](comparator functions.Comparator[T]) streams.Collector[T, *Queue[T]] {
	return sequences.Collector[T, *Queue[T]]{New[T](comparator)}
}

func (pq *Queue[T]) Empty() bool {
	return len(pq.slice) == 0
}

func (pq *Queue[T]) PushBack(value T) {
	pq.slice = append(pq.slice, value)

	currNode := len(pq.slice) - 1
	for {
		parentNode := (currNode - 1) / 2
		if parentNode == currNode || pq.comparator(pq.slice[currNode], pq.slice[parentNode]) != comparison.FirstSmaller {
			break
		}

		pq.slice[parentNode], pq.slice[currNode] = pq.slice[currNode], pq.slice[parentNode]
		currNode = parentNode
	}
}

func (pq *Queue[T]) PeekFront() T {
	if pq.Empty() {
		panic("Priority queue is empty")
	}

	return pq.slice[0]
}

func (pq *Queue[T]) PopFront() T {
	if pq.Empty() {
		panic("Priority queue is empty")
	}

	lastIndex := len(pq.slice) - 1

	pq.slice[0], pq.slice[lastIndex] = pq.slice[lastIndex], pq.slice[0]

	currNode := 0
	for {
		leftChild := 2*currNode + 1
		if leftChild >= lastIndex || leftChild < 0 {
			break
		}

		child := leftChild
		if rightChild := leftChild + 1; rightChild < lastIndex && pq.comparator(pq.slice[rightChild], pq.slice[leftChild]) == comparison.FirstSmaller {
			child = rightChild
		}

		if pq.comparator(pq.slice[child], pq.slice[currNode]) != comparison.FirstSmaller {
			break
		}

		pq.slice[currNode], pq.slice[child] = pq.slice[child], pq.slice[currNode]
		currNode = child
	}

	lastElem := pq.slice[lastIndex]
	pq.slice = pq.slice[:lastIndex]
	return lastElem
}
