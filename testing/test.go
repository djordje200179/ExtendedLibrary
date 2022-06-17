package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/other/priorityqueue"
)

func main() {
	queue := priorityqueue.New[int](priorityqueue.BiggerFirst)

	queue.Push(1, 1)
	queue.Push(3, 3)
	queue.Push(5, 5)
	queue.Push(2, 2)
	queue.Push(0, 0)
	queue.Push(7, 7)

	queue.ForEach(func(value int) {
		fmt.Println(value)
	})
}
