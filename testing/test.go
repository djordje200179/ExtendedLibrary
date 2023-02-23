package main

import (
	"github.com/djordje200179/extendedlibrary/datastructures/maps/binarysearchtree"
	"github.com/djordje200179/extendedlibrary/misc/functions/comparison"
)

func main() {
	tree := binarysearchtree.New[int, string](comparison.Ascending[int])

	tree.Set(1, "one")
	tree.Set(5, "five")
	tree.Set(3, "three")

	for node := tree.M
}
