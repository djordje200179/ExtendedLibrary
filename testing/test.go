package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/linkedlistmap"
	"github.com/djordje200179/extendedlibrary/misc"
)

func main() {
	m := linkedlistmap.New[string, int]()

	m.Set("abc", 2)
	m.Set("abcd", 5)

	m.Stream().ForEach(func(pair misc.Pair[string, int]) {
		fmt.Println(pair.First, pair.Second)
	})
}
