package iterable

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type IteratorStream[T any] struct {
	Iterator[T]
}

func (stream IteratorStream[T]) Supply() optional.Optional[T] {
	if stream.Valid() {
		defer stream.Move()
		return optional.FromValue(stream.Get())
	} else {
		return optional.Empty[T]()
	}
}
