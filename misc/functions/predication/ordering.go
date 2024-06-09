package predication

import (
	"cmp"
)

// LessThan returns a Predicate that checks
// if the value is less than the parameter.
func LessThan[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool { return value < param }
}

// LessThanOrEqual returns a Predicate that checks
// if the value is less than or equal to the parameter.
func LessThanOrEqual[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool { return value <= param }
}

// GreaterThan returns a Predicate that checks
// if the value is greater than the parameter.
func GreaterThan[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool { return value > param }
}

// GreaterThanOrEqual returns a Predicate that checks
// if the value is greater than or equal to the parameter.
func GreaterThanOrEqual[T cmp.Ordered](param T) Predicate[T] {
	return func(value T) bool { return value >= param }
}
