package suppliers

import (
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type Channel[T any] <-chan T

func (supplier Channel[T]) Supply() optional.Optional[T] {
	data, ok := <-supplier

	return optional.Optional[T]{
		Value: data,
		Valid: ok,
	}
}
