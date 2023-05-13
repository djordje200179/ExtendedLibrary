package predication

func Equals[T comparable](param T) Predicate[T] {
	return func(value T) bool {
		return value == param
	}
}

func NotEquals[T comparable](param T) Predicate[T] {
	return func(value T) bool {
		return value != param
	}
}

func IsNil[T any]() Predicate[*T] {
	return func(value *T) bool {
		return value == nil
	}
}

func IsNotNil[T any]() Predicate[*T] {
	return func(value *T) bool {
		return value != nil
	}
}
