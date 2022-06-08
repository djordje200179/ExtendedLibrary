package main

import (
	"fmt"
	"github.com/djordje200179/GoExtendedLibrary/misc/functions"
	"github.com/djordje200179/GoExtendedLibrary/streams"
)

func IsPrime(num int) bool {
	return streams.Range(2, num/2+1).All(func(val int) bool { return num%val != 0 })
}

func main() {
	streams.Generate(2, functions.Increment).Filter(IsPrime).ForEach(func(val int) { fmt.Println(val) })
}
