package optional

import (
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

type Optional[T any] struct {
	Value T
	Valid bool
}

func Empty[T any]() Optional[T] {
	return Optional[T]{
		Valid: false,
	}
}

func FromValue[T any](value T) Optional[T] {
	return Optional[T]{value, true}
}

func (o Optional[T]) GetOrElse(other T) T {
	if o.Valid {
		return o.Value
	} else {
		return other
	}
}

func (o Optional[T]) GetOrDefault() T {
	var def T
	return o.GetOrElse(def)
}

func (o Optional[T]) Process(onValue functions.ParamCallback[T], onEmpty functions.EmptyCallback) {
	if o.Valid {
		onValue(o.Value)
	} else {
		onEmpty()
	}
}

func Map[T any, P any](o Optional[T], mapper functions.Mapper[T, P]) Optional[P] {
	if o.Valid {
		return FromValue(mapper(o.Value))
	} else {
		return Empty[P]()
	}
}

func (o Optional[T]) Filter(predicate predication.Predictor[T]) Optional[T] {
	if o.Valid && predicate(o.Value) {
		return o
	} else {
		return Empty[T]()
	}
}
