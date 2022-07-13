package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type channelSupplier[T any] struct {
	channel <-chan T
}

func FromChannel[T any](channel <-chan T) Supplier[T] { return &channelSupplier[T]{channel} }

func (supplier channelSupplier[T]) Supply() optional.Optional[T] {
	data, ok := <-supplier.channel
	return optional.New(data, ok)
}
