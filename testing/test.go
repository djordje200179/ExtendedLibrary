package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/linkedlistmap"
)

func main() {
	m := linkedlistmap.New[int, int]()
	m.Set(10, 1)
	m.Set(5, 2)
	m.Set(7, 3)
	m.Set(11, 4)

	for it := m.Iterator(); it.Valid(); it.Move() {
		entry := it.Get()
		fmt.Println(entry.Key(), " ", entry.Value())
	}
}
