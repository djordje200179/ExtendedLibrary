package messenger

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/optional"
	"unsafe"
)

type Messenger[T any] chan T

func New[T any](bufferSize int) Messenger[T] {
	return make(chan T, bufferSize)
}

func (messenger Messenger[T]) Send(signal T) {
	messenger <- signal
}

func (messenger Messenger[T]) ReadSync() optional.Optional[T] {
	if messenger.Closed() {
		return optional.Empty[T]()
	} else {
		return optional.FromValue(<-messenger)
	}
}

func (messenger Messenger[T]) ReadAsync(callback functions.ParamCallback[T]) {
	go func() {
		if messenger.Closed() {
			return
		}

		value := <-messenger
		callback(value)
	}()
}

func (messenger Messenger[T]) ReadIfHasData() optional.Optional[T] {
	if messenger.Closed() {
		return optional.Empty[T]()
	}

	select {
	case data := <-messenger:
		return optional.FromValue(data)
	default:
		return optional.Empty[T]()
	}
}

func (messenger Messenger[T]) Closed() bool {
	cptr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&messenger)) + unsafe.Sizeof(uint(0))))
	cptr += unsafe.Sizeof(uint(0)) * 2
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	cptr += unsafe.Sizeof(uint16(0))
	return *(*uint32)(unsafe.Pointer(cptr)) > 0
}

func (messenger Messenger[T]) Close() {
	close(messenger)
}
