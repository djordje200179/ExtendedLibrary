package iterable

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

// IteratorStream creates a streams.Stream from an Iterator.
func IteratorStream[T any](iterator Iterator[T]) streams.Stream[T] {
	supplier := func() optional.Optional[T] {
		if iterator.Valid() {
			defer iterator.Move()
			return optional.FromValue(iterator.Get())
		} else {
			return optional.Empty[T]()
		}
	}

	return streams.New(supplier)
}
