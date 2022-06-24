package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
)

func main() {
	l := linkedlist.New[int]()
	l.Append(2)
	l.AppendMany(3, 5)

	l.RefStream().ForEach(func(ref *int) {
		fmt.Println(*ref)
	})
}
