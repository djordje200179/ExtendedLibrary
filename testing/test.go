package main

import "github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"

func main() {
	list := linkedlist.New[int]()
	list.Stream().ForEach(func(item int) { println(item) })
}
