package main

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

func main() {
	streams.
		Generate(1, functions.Increment).
		Filter(func(value int) bool { return value%6 == 0 }).
		Seek(5).
		Limit(10).
		ForEach(func(value int) { fmt.Println(value) })
}
