package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"github.com/djordje200179/extendedlibrary/streams"
)

func Channel[T any](channel <-chan T) streams.Supplier[T] {
	return func() optional.Optional[T] {
		data, ok := <-channel

		return optional.Optional[T]{
			Value: data,
			Valid: ok,
		}
	}
}
