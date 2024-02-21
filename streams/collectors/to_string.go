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

func (c *stringCollector[T]) Supply(value T) {
	if c.builder.Len() > 0 {
		c.builder.WriteString(c.delimiter)
	}

	str := fmt.Sprint(value)
	c.builder.WriteString(str)
}

func (c *stringCollector[T]) Finish() string {
	return c.builder.String()
}
