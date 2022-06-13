package functions

type ParamCallback[T any] func(value T)
type EmptyCallback func()

type Predicate[T any] func(value T) bool

type Generator[T any] func() T
