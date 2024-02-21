package predication

import "github.com/djordje200179/extendedlibrary/misc/functions"

type Predicate[T any] functions.Mapper[T, bool]

func (p Predicate[T]) And(other Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return p(value) && other(value)
	}
}

func (p Predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return p(value) || other(value)
	}
}

func (p Predicate[T]) Not() Predicate[T] {
	return func(value T) bool {
		return !p(value)
	}
}
