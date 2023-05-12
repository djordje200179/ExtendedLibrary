package functions

type ParamCallback[T any] func(value T)
type EmptyCallback func()

type Mapper[T, P any] func(value T) P
