package predication

import (
	"cmp"
)

func LessThan[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value < param
	}
}

func LessThanOrEqual[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value <= param
	}
}

func GreaterThan[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value > param
	}
}

func GreaterThanOrEqual[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool {
		return value >= param
	}
}
