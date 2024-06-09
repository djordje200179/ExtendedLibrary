package predication

// Equals returns a Predicate that checks
// if the value is equal to the parameter.
func Equals[T comparable](param T) Predicate[T] {
	return func(value T) bool { return value == param }
}

// NotEquals returns a Predicate that checks
// if the value is not equal to the parameter.
func NotEquals[T comparable](param T) Predicate[T] {
	return func(value T) bool { return value != param }
}

// IsNil returns a Predicate that checks if the value is nil.
func IsNil[T any]() Predicate[*T] {
	return func(value *T) bool { return value == nil }
}

// IsNotNil returns a Predicate that checks if the value is not nil.
func IsNotNil[T any]() Predicate[*T] {
	return func(value *T) bool { return value != nil }
}
