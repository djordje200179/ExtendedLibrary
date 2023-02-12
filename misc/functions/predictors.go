package functions

import "golang.org/x/exp/constraints"

func Equals[T comparable](param T) Predictor[T] {
	return func(value T) bool {
		return value == param
	}
}

func NotEquals[T comparable](param T) Predictor[T] {
	return func(value T) bool {
		return value != param
	}
}

func LessThan[T constraints.Ordered](param T) Predictor[T] {
	return func(value T) bool {
		return value < param
	}
}

func LessThanOrEqual[T constraints.Ordered](param T) Predictor[T] {
	return func(value T) bool {
		return value <= param
	}
}

func GreaterThan[T constraints.Ordered](param T) Predictor[T] {
	return func(value T) bool {
		return value > param
	}
}

func GreaterThanOrEqual[T constraints.Ordered](param T) Predictor[T] {
	return func(value T) bool {
		return value >= param
	}
}

func IsNil[T any]() Predictor[*T] {
	return func(value *T) bool {
		return value == nil
	}
}

func IsNotNil[T any]() Predictor[*T] {
	return func(value *T) bool {
		return value != nil
	}
}
