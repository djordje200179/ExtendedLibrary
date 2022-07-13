package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/streams"
	"github.com/djordje200179/extendedlibrary/streams/collectors"
	"math"
)

func countDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)) + 1))
}
func main() {
	fmt.Print(streams.Collect(streams.FromRange(1, 1001), collectors.Group(countDigits)))
}
