package main

import (
	"github.com/djordje200179/extendedlibrary/datastructures/set"
)

func main() {
	s := set.NewHashset[int]()

	s.Add(1)
	s.Add(6)
	s.Add(2)
	s.Add(5)
	s.Add(1)
	s.Add(5)
	s.Add(10)
	s.Add(1)

	for it := s.Iterator(); it.Valid(); it.Move() {
		println(it.Get())
	}
}
