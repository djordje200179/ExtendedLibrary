package predication

import "github.com/djordje200179/extendedlibrary/misc/functions"

// Predicate is a function that takes a value and returns a boolean.
type Predicate[T any] functions.Mapper[T, bool]

// And creates a new Predicate that returns true
// if both the original Predicate and the other Predicate return true.
func (p Predicate[T]) And(other Predicate[T]) Predicate[T] {
	return func(value T) bool { return p(value) && other(value) }
}

// Or creates a new Predicate that returns true
// if either the original Predicate or the other Predicate return true.
func (p Predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return func(value T) bool { return p(value) || other(value) }
}

// Not creates a new Predicate that returns
// the opposite of the original Predicate.
func (p Predicate[T]) Not() Predicate[T] {
	return func(value T) bool { return !p(value) }
}
