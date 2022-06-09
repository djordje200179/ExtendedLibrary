package functions

type Less[T any] func(first, second T) bool

type ParamCallback[T any] func(value T)
type EmptyCallback func()

type Predicate[T any] func(value T) bool
