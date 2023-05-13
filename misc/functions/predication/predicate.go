package predication

import "github.com/djordje200179/extendedlibrary/misc/functions"

type Predicate[T any] functions.Mapper[T, bool]

func (predicate Predicate[T]) And(other Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return predicate(value) && other(value)
	}
}

func (predicate Predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return func(value T) bool {
		return predicate(value) || other(value)
	}
}

func (predicate Predicate[T]) Not() Predicate[T] {
	return func(value T) bool {
		return !predicate(value)
	}
}
