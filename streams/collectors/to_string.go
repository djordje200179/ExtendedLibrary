package collectors

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/streams"
	"strings"
)

type stringCollector[T any] struct {
	builder   strings.Builder
	delimiter string
}

func ToString[T any](delimiter string) streams.Collector[T, string] {
	return &stringCollector[T]{
		builder:   strings.Builder{},
		delimiter: delimiter,
	}
}

func (collector *stringCollector[T]) Supply(value T) {
	if collector.builder.Len() > 0 {
		collector.builder.WriteString(collector.delimiter)
	}

	str := fmt.Sprint(value)
	collector.builder.WriteString(str)
}

func (collector *stringCollector[T]) Finish() string {
	return collector.builder.String()
}
