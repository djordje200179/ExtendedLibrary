package predication

import (
	"golang.org/x/exp/constraints"
)

func LessThan[T constraints.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value < param
	}
}

func LessThanOrEqual[T constraints.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value <= param
	}
}

func GreaterThan[T constraints.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value > param
	}
}

func GreaterThanOrEqual[T constraints.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value >= param
	}
}
