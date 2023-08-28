package cols

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

// RefsStream returns a streams.Stream of references to the elements supplied by the given Iterator.
func RefsStream[T any](iterator Iterator[T]) streams.Stream[*T] {
	supplier := func() optional.Optional[*T] {
		if iterator.Valid() {
			defer iterator.Move()
			return optional.FromValue(iterator.GetRef())
		} else {
			return optional.Empty[*T]()
		}
	}

	return streams.New(supplier)
}
