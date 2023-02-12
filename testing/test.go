package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/datastructures/sets/mapset"
)

func main() {
	set := mapset.NewHashSet[string]()

	set.Add("Macka")
	set.Add("Pas")
	set.Add("Konj")
	set.Add("Konj")
	set.Add("Konj")

	for it := set.Iterator(); it.Valid(); it.Move() {
		fmt.Println(it.Get())
	}
}
