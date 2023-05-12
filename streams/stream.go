package streams

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type Supplier[T any] func() optional.Optional[T]

type Stream[T any] struct {
	supplier Supplier[T]
}

type Streamer[T any] interface {
	Stream() Stream[T]
}

func New[T any](supplier Supplier[T]) Stream[T] {
	return Stream[T]{supplier}
}
