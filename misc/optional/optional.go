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

func (o Optional[T]) HasValue() bool {
	return o.valid
}

func (o Optional[T]) GetPair() (T, bool) {
	return o.value, o.valid
}

func (o Optional[T]) Get() T {
	if o.valid {
		return o.value
	} else {
		panic("No value present")
	}
}

func (o Optional[T]) GetOrElse(other T) T {
	if o.valid {
		return o.value
	} else {
		return other
	}
}

func (o Optional[T]) GetOrDefault() T {
	var def T
	return o.GetOrElse(def)
}

func (o Optional[T]) Process(onValue functions.ParamCallback[T], onEmpty functions.EmptyCallback) {
	if o.valid {
		onValue(o.value)
	} else {
		onEmpty()
	}
}

func Map[T any, P any](o Optional[T], mapper functions.Mapper[T, P]) Optional[P] {
	if o.valid {
		return FromValue(mapper(o.value))
	} else {
		return Empty[P]()
	}
}

func (o Optional[T]) Filter(predicate functions.Predictor[T]) Optional[T] {
	if o.valid && predicate(o.value) {
		return o
	} else {
		return Empty[T]()
	}
}
