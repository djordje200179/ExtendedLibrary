package predication

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
