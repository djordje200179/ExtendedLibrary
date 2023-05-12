package functions

func Identity[T any](value T) T {
	return value
}

func Zero[T any]() T {
	var zero T
	return zero
}
