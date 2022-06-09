package optional

import "github.com/djordje200179/extendedlibrary/misc/functions"

type Optional[T any] struct {
	value T
	valid bool
}

func Empty[T any]() Optional[T] {
	return Optional[T]{
		valid: false,
	}
}

func FromValue[T any](value T) Optional[T] {
	return New[T](value, true)
}

func New[T any](value T, valid bool) Optional[T] {
	return Optional[T]{
		value: value,
		valid: valid,
	}
}

func (o Optional[T]) IsPresent() bool {
	return o.valid
}

func (o Optional[T]) Get() (T, bool) {
	return o.value, o.valid
}

func (o Optional[T]) GetOrElse(other T) T {
	if o.valid {
		return o.value
	} else {
		return other
	}
}

func (o Optional[T]) GetOrPanic() T {
	if o.valid {
		return o.value
	} else {
		panic("No value present")
	}
}

func (o Optional[T]) Process(onValue functions.ParamCallback[T], onEmpty functions.EmptyCallback) {
	if o.valid {
		onValue(o.value)
	} else {
		onEmpty()
	}
}
