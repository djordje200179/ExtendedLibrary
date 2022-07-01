package boundedbuffer

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
)

type BoundedBuffer[T any] chan T

func New[T any](bufferSize int) BoundedBuffer[T] { return make(chan T, bufferSize) }

func (buffer BoundedBuffer[T]) Put(signal T) { buffer <- signal }

func (buffer BoundedBuffer[T]) Get() optional.Optional[T] {
	value, ok := <-buffer
	return optional.New[T](value, ok)
}

func (buffer BoundedBuffer[T]) GetAsync(callback functions.ParamCallback[T]) {
	go func() {
		value, ok := <-buffer

		if ok {
			callback(value)
		}
	}()
}

func (buffer BoundedBuffer[T]) ReadImmediately() optional.Optional[T] {
	select {
	case data := <-buffer:
		return optional.FromValue(data)
	default:
		return optional.Empty[T]()
	}
}

func (buffer BoundedBuffer[T]) Close() { close(buffer) }
