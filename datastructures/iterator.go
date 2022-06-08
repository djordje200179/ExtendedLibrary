package datastructures

type Iterator[T any] interface {
	IsValid() bool
	Move()

	Get() T
}
