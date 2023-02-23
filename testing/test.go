package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/binarysearchtree"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

func main() {
	tree := binarysearchtree.New[int, string](comparison.Ascending[int])

	tree.Set(1, "one")
	tree.Set(5, "five")
	tree.Set(3, "three")
	tree.Set(2, "four")
	tree.Set(3, "three-2")

	for node := tree.Root().Min(); node != nil; node = node.Next() {
		fmt.Println(node.Key(), node.Value)
	}
}
