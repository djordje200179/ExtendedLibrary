package functions

type Predictor[T any] Mapper[T, bool]

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
