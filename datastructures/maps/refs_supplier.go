package maps

import (
	"github.com/djordje200179/extendedlibrary/misc"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

func RefsStream[K, V any](m Map[K, V]) streams.Stream[misc.Pair[K, *V]] {
	iterator := m.ModifyingIterator()
	supplier := func() optional.Optional[misc.Pair[K, *V]] {
		if !iterator.Valid() {
			return optional.Empty[misc.Pair[K, *V]]()
		}

		pair := misc.MakePair(iterator.Key(), iterator.ValueRef())
		iterator.Move()

		return optional.FromValue(pair)
	}

	return streams.New(supplier)
}
