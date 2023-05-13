package optional

import (
	"fmt"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/misc/functions/predication"
)

type Optional[T any] struct {
	Value T
	Valid bool
}

func New[T any](value T, valid bool) Optional[T] {
	return Optional[T]{value, valid}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{
		Valid: false,
	}
}

func FromValue[T any](value T) Optional[T] {
	return Optional[T]{value, true}
}

func (optional Optional[T]) GetOrElse(other T) T {
	if optional.Valid {
		return optional.Value
	} else {
		return other
	}
}

func (optional Optional[T]) GetOrDefault() T {
	var def T
	return optional.GetOrElse(def)
}

func (optional Optional[T]) Process(onValue functions.ParamCallback[T], onEmpty functions.EmptyCallback) {
	if optional.Valid {
		onValue(optional.Value)
	} else {
		onEmpty()
	}
}

func Map[T any, P any](optional Optional[T], mapper functions.Mapper[T, P]) Optional[P] {
	if optional.Valid {
		return FromValue(mapper(optional.Value))
	} else {
		return Empty[P]()
	}
}

func (optional Optional[T]) Filter(predicate predication.Predictor[T]) Optional[T] {
	if optional.Valid && predicate(optional.Value) {
		return optional
	} else {
		return Empty[T]()
	}
}

func (optional Optional[T]) String() string {
	if optional.Valid {
		return fmt.Sprint(optional.Value)
	} else {
		return "(empty)"
	}
}
