package predictors

import "github.com/djordje200179/extendedlibrary/misc/functions"

func Equals[T comparable](param T) functions.Predictor[T] {
	return func(value T) bool {
		return value == param
	}
}

func NotEquals[T comparable](param T) functions.Predictor[T] {
	return func(value T) bool {
		return value != param
	}
}

func IsNil[T any]() functions.Predictor[*T] {
	return func(value *T) bool {
		return value == nil
	}
}

func IsNotNil[T any]() functions.Predictor[*T] {
	return func(value *T) bool {
		return value != nil
	}
}
