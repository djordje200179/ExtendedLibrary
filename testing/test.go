package main

import (
	"fmt"
	streams2 "github.com/djordje200179/extendedlibrary/datastructures/extensions/streams"
	"github.com/djordje200179/extendedlibrary/datastructures/maps/hashmap"
	"github.com/djordje200179/extendedlibrary/datastructures/sequences/linkedlist"
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/collectors"
	"math"
	"math/rand"
)

func countDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)) + 1))
}

func countElements(pair misc.Pair[int, []int]) misc.Pair[int, int] {
	return misc.Pair[int, int]{pair.First, len(pair.Second)}
}

func main() {
	nums := linkedlist.New[int]()
	for i := 0; i < 10000; i++ {
		nums.Append(rand.Int())
	}

	dict := hashmap.NewFromMap(streams.Collect(streams2.FromIterable[int](nums), collectors.Group[int, int](countDigits)))
	counters := streams.Map(streams2.FromMap[int, []int](dict), countElements)

	for it := streams2.Iterator(counters); it.Valid(); it.Move() {
		entry := it.Get()
		fmt.Printf("%v: %v\n", entry.First, entry.Second)
	}
}
