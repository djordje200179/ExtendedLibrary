package misc

type Cloner[T any] interface {
	Clone() T
}
