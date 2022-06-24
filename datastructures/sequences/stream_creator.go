package sequences

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

func RefStream[T any](sequence Sequence[T]) *streams.Stream[*T] {
	it := sequence.ModifyingIterator()
	supplier := func() optional.Optional[*T] {
		if !it.Valid() {
			return optional.Empty[*T]()
		}

		defer it.Move()
		return optional.FromValue(it.GetRef())
	}

	return streams.SupplyWithEnd[*T](supplier)
}
