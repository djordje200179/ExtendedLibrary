package functions

// Mapper is a function type that maps a value of type T to a value of type P.
type Mapper[T, P any] func(value T) P
