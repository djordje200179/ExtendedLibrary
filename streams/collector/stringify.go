package collector

import (
	"fmt"
	"strings"
)

type stringify[T any] struct {
	builder   *strings.Builder
	delimiter string
}

func Stringify[T any](delimiter string) Collector[T, string] {
	return stringify[T]{
		builder:   new(strings.Builder),
		delimiter: delimiter,
	}
}

func (stringify stringify[T]) Supply(value T) {
	if stringify.builder.Len() > 0 {
		stringify.builder.WriteString(stringify.delimiter)
	}

	str := fmt.Sprint(value)
	stringify.builder.WriteString(str)
}

func (stringify stringify[T]) Finish() string {
	return stringify.builder.String()
}
