package main

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/heapqueue"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

func main() {
	pq := heapqueue.New[int](comparison.Ascending[int])

	pq.PushBack(1)
	pq.PushBack(8)
	pq.PushBack(3)
	pq.PushBack(5)
	pq.PushBack(8)
	pq.PushBack(2)

	for !pq.Empty() {
		println(pq.PopFront())
	}
}
