package messenger

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type Messenger[T any] chan T

func New[T any](bufferSize int) Messenger[T] { return make(chan T, bufferSize) }

func (messenger Messenger[T]) Send(signal T) { messenger <- signal }

func (messenger Messenger[T]) ReadSync() optional.Optional[T] {
	value, ok := <-messenger
	return optional.New[T](value, ok)
}

func (messenger Messenger[T]) ReadAsync(callback functions.ParamCallback[T]) {
	go func() {
		value, ok := <-messenger

		if ok {
			callback(value)
		}
	}()
}

func (messenger Messenger[T]) ReadIfHasData() optional.Optional[T] {
	select {
	case data := <-messenger:
		return optional.FromValue(data)
	default:
		return optional.Empty[T]()
	}
}

func (messenger Messenger[T]) Close() { close(messenger) }
