package sources

import (
	"github.com/djordje200179/extendedlibrary/mapreduce"
	"github.com/djordje200179/extendedlibrary/misc"
)

func NewSliceSource[T any](slice []T) mapreduce.Source[int, T] {
	source := make(chan misc.Pair[int, T], 100)

	go func() {
		for index, element := range slice {
			source <- misc.Pair[int, T]{index, element}
		}
		close(source)
	}()

	return source
}
