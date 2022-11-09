package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type Channel[T any] struct {
	Channel <-chan T
}

func (supplier Channel[T]) Supply() optional.Optional[T] {
	data, ok := <-supplier.Channel
	return optional.Optional[T]{data, ok}
}
