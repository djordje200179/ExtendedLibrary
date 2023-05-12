package predication

import "github.com/djordje200179/extendedlibrary/misc/functions"

type Predictor[T any] functions.Mapper[T, bool]

func (p Predictor[T]) And(other Predictor[T]) Predictor[T] {
	return func(value T) bool {
		return p(value) && other(value)
	}
}

func (p Predictor[T]) Or(other Predictor[T]) Predictor[T] {
	return func(value T) bool {
		return p(value) || other(value)
	}
}

func (p Predictor[T]) Not() Predictor[T] {
	return func(value T) bool {
		return !p(value)
	}
}
